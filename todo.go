package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "todo"
	app.Version = Version
	app.Usage = ""
	app.Author = "tcnksm"
	app.Email = "nsd22843@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
