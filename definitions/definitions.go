package definitions

import (
	"bytes"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/hashicorp/hcl"
)

const (
	goTemplate = `
package cmd

func init() {
{{ .Commands }}
}

func root(c *cobra.Command, cmd string) {

{{ .MegaSwitch }} 
}
`
	switchTemplate = `switch cmd:
{{ range .SwitchStatements }} 
case "{{.CommandName}}":
  api.{{.CommandName}}


`
	cmdTemplate = `var {{.VariableName}} = &cobra.Command{
  Use:   "{{- .Name -}}",
  Short: "{{- .ShortDescription -}}",
  Long: ` + "`" + `{{- .Description -}}` + "`" + `,
{{ if .V4APIName}}Run: func(cmd *cobra.Command, args []string) {
    root(cmd, "{{- .V4APIName}}")
  },{{ end}}
}
{{- $varName := .VariableName }}
{{range .Subcommands}}{{$varName}}.AddCommand({{.}}){{end}}
{{ if .TopLevel }}RootCmd.AddCommand({{.VariableName}}){{ end }}
`
)

type Command struct {
	Name             string
	Description      string
	ShortDescription string
	V4APIName        string
	Options          map[string]string
	Subcommands      []string
	TopLevel         bool
}

type CommandTemplateValues struct {
	Name             string
	VariableName     string
	V4APIName        string
	Description      string
	ShortDescription string

	Subcommands []string
	TopLevel    bool
}

type FileTemplateValues struct {
	Commands   string
	MegaSwitch string
}

func fileText(commands, switches string) (string, error) {
	var (
		buff bytes.Buffer
	)
	tmpl, err := template.New("gofile").Parse(goTemplate)
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(&buff, &FileTemplateValues{
		Commands:   commands,
		MegaSwitch: switches,
	})
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
	})
	if err != nil {
		return "", err
	}
	return buff.String(), nil
}

func LoadDefinitions(fname string) ([]*Command, error) {
	c := []*Command{}
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	err = hcl.Decode(&c, string(buf))
	if err != nil {
		return nil, err
	}
	return c, nil
}

func hyphenDelimToCamel(s string) string {
	parts := strings.Split(s, "-")
	if len(parts) == 1 {
		return s
	}

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
	for _, cmd := range cmds {
		if cmd.TopLevel {
			s, err := cmd.ToGo()
			if err != nil {
				return err
			}
			commandsGo += s
		}
	}

	t, err := fileText(commandsGo, "")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(outfile, []byte(t), 0644)
	if err != nil {
		return err
	}
	return nil
}
