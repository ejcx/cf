package definitions

import (
	"bytes"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/BurntSushi/toml"
)

const (
	goTemplate = `
package cmd

var (
{{ .Variables }}
)

func init() {
{{ .Commands }}
}
`

	varTemplate = `{{ range .}}{{ .Name }} {{ .Type }}
{{ end }}`

	cmdTemplate = `var {{.VariableName}} = &cobra.Command{
  Use:   "{{- .Name -}}",
  Short: "{{- .ShortDescription -}}",
  Long: ` + "`" + `{{- .Description -}}` + "`" + `,
{{ if .V4APIName}}Run: func(cmd *cobra.Command, args []string) {
    Main(cmd, args, "{{- .V4APIName}}")
  },{{ end}}
}
{{- $varName := .VariableName }}
{{range .Subcommands}}{{$varName}}.AddCommand({{.}})
{{end}}
{{ if .TopLevel }}RootCmd.AddCommand({{.VariableName}})
{{ end }}
{{range .Option}}
  {{$varName}}.Flags().{{.TypeCap}}Var(&{{.ArgName}}, "{{.Name}}", {{.Default}}, "{{.Description}}")
  {{if .Required}}{{$varName}}.MarkFlagRequired("{{.Name}}")
{{end}}
{{end}}
`
)

var (
	declaredVariables = make(map[string]bool)
)

type Option struct {
	Name        string
	Type        string
	Description string
	Required    bool
}

type OptionTemplateValue struct {
	TypeCap     string
	ArgName     string
	Name        string
	Default     string
	Description string
	Required    bool
}

type Command struct {
	Name             string
	Description      string
	ShortDescription string
	V4APIName        string
	Option           []Option
	Subcommands      []string
	TopLevel         bool
}

type CommandTemplateValues struct {
	Name             string
	VariableName     string
	V4APIName        string
	Description      string
	ShortDescription string
	Option           []OptionTemplateValue

	Subcommands []string
	TopLevel    bool
}

type FileTemplateValues struct {
	Commands  string
	Variables string
}

func fileText(commands, variables string) (string, error) {
	var (
		buff bytes.Buffer
	)
	tmpl, err := template.New("gofile").Parse(goTemplate)
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(&buff, &FileTemplateValues{
		Commands:  commands,
		Variables: variables,
	})
	if err != nil {
		return "", err
	}
	return buff.String(), nil
}

func (o Option) ToOptionTemplateValue() OptionTemplateValue {
	defaultType := "\"\""
	if o.Type == "int" {
		defaultType = "0"
	} else if o.Type == "bool" {
		defaultType = "false"
	}
	return OptionTemplateValue{
		TypeCap:     hyphenDelimToCamel(o.Type),
		ArgName:     hyphenDelimToCamel(o.Name),
		Name:        o.Name,
		Default:     defaultType,
		Description: o.Description,
		Required:    o.Required,
	}
}

func (c *Command) ToVariables() (string, error) {
	var (
		buff    bytes.Buffer
		options []Option
	)

	// We need to do a pass to convert the variable names from cmd flags
	// to Go variable names.
	for _, opt := range c.Option {
		if _, ok := declaredVariables[opt.Name]; ok {
			continue
		}
		options = append(options, Option{
			Name: hyphenDelimToCamel(opt.Name),
			Type: opt.Type,
		})
		declaredVariables[opt.Name] = true
	}

	tmpl, err := template.New("option").Parse(varTemplate)
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(&buff, options)
	if err != nil {
		return "", err
	}
	return buff.String(), nil
}
func (c *Command) ToGo() (string, error) {
	var (
		buff               bytes.Buffer
		subcommandVarNames []string
	)
	for _, subcmd := range c.Subcommands {
		subcommandVarNames = append(subcommandVarNames, hyphenDelimToCamel(subcmd))
	}

	// Convert the options to OptionTemplateValues
	var optionTemplateValueList []OptionTemplateValue
	for _, opt := range c.Option {
		optionTemplateValueList = append(optionTemplateValueList, opt.ToOptionTemplateValue())
	}

	tmpl, err := template.New("command").Parse(cmdTemplate)
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(&buff, &CommandTemplateValues{
		Name:             c.Name,
		TopLevel:         c.TopLevel,
		Subcommands:      subcommandVarNames,
		VariableName:     hyphenDelimToCamel(c.Name),
		V4APIName:        c.V4APIName,
		ShortDescription: c.ShortDescription,
		Description:      c.Description,
		Option:           optionTemplateValueList,
	})
	if err != nil {
		return "", err
	}
	return buff.String(), nil
}

func LoadDefinitions(fname string) ([]*Command, error) {
	c := map[string][]*Command{}
	_, err := toml.DecodeFile(fname, &c)
	if err != nil {
		return nil, err
	}
	return c["command"], nil
}

func hyphenDelimToCamel(s string) string {
	parts := strings.Split(s, "-")

	// Capitalize the first letter of each part and return the
	// concatenation of them.
	for i, part := range parts {
		parts[i] = strings.ToUpper(string(part[0])) + part[1:]
	}
	return strings.Join(parts, "")
}

func GenerateFile(fname string, outfile string) error {
	cmds, err := LoadDefinitions(fname)
	if err != nil {
		return err
	}
	commandsGo := ""
	variablesGo := ""
	// Do two passes. First we want to generate all code that is
	// not a top level command. Next we want to generate all top
	// level commands. This is because of the variable ordering
	// and avoiding undefined variable issues.
	for _, cmd := range cmds {
		if !cmd.TopLevel {
			s, err := cmd.ToGo()
			if err != nil {
				return err
			}
			commandsGo += s
		}
	}

	// While we do this second pass, create all the variables while
	// we are also creating all of the commands.
	for _, cmd := range cmds {
		v, err := cmd.ToVariables()
		if err != nil {
			return err
		}
		variablesGo += v
		if cmd.TopLevel {
			s, err := cmd.ToGo()
			if err != nil {
				return err
			}
			commandsGo += s
		}
	}

	t, err := fileText(commandsGo, variablesGo)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(outfile, []byte(t), 0644)
	if err != nil {
		return err
	}
	return nil
}
