package main

import "github.com/alexandrst88/terraform-variables-generator/cmd"

// Version is updated by linker flags during build time
var Version = "dev"

func main() {
	cmd.Execute(Version)
}
