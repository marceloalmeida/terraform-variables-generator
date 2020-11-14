package main

import "github.com/marceloalmeida/terraform-variables-generator/cmd"

// Version is updated by linker flags during build time
var Version = "dev"

func main() {
	cmd.Execute(Version)
}
