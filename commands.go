package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var Commands = []cli.Command{
	commandAdd,
	commandList,
	commandDelete,
}

var commandAdd = cli.Command{
	Name:  "add",
	Usage: "",
	Description: `
`,
	Action: doAdd,
}

var commandList = cli.Command{
	Name:  "list",
	Usage: "",
	Description: `
`,
	Action: doList,
}

var commandDelete = cli.Command{
	Name:  "delete",
	Usage: "",
	Description: `
`,
	Action: doDelete,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doAdd(c *cli.Context) {
}

func doList(c *cli.Context) {
}

func doDelete(c *cli.Context) {
}
