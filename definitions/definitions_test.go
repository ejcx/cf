package definitions

import (
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	ast, err := LoadDefinitions()
	if err != nil {
		t.Fatalf("Could not load configuration file: %s", err)
	}
	if ast[0].Subcommands == "" {
		t.Fatalf("Subcommand value is empty")
	}
}
