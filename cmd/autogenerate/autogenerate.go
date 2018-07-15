package main

import (
	"log"
	"os/exec"

	"github.com/ejcx/cf/definitions"
)

func main() {
	err := definitions.GenerateFile("definitions/definitions.toml", "cmd/autogenerated.go")
	if err != nil {
		log.Fatalf("Could not generate file: %s", err)
	}
	err = exec.Command("goimports", "-w", "cmd/autogenerated.go").Run()
	if err != nil {
		log.Fatalf("Could not run goimports on autogenerated file: %s", err)
	}
	err = exec.Command("go", "fmt", "cmd/autogenerated.go").Run()
	if err != nil {
		log.Fatalf("Could not run go fmt on autogenerated file: %s", err)
	}

}
