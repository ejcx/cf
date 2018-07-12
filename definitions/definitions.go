package definitions

import (
	"io/ioutil"

	"github.com/hashicorp/hcl"
)

type Command struct {
	Name        string
	Description string
	V4APIName   string
	Options     map[string]string
	Subcommands []string
}

func LoadDefinitions() ([]*Command, error) {
	c := []*Command{}
	buf, err := ioutil.ReadFile("definitions.hcl")
	if err != nil {
		return nil, err
	}
	// ast, err := hcl.Parse(string(buf))
	// if err != nil {
	// 	return nil, err
	// }
	err = hcl.Decode(&c, string(buf))
	if err != nil {
		return nil, err
	}
	return c, nil
}

func PrintCobraCommand() {

}
