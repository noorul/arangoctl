package main

import (
	"github.com/galaxyse/arangoctl/cmd/arangoctl/subcmd"
)

var (
	Version = "dev"
)

func main() {
	subcmd.Execute(Version)
}
