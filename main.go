package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "myprogramname"
	app.Commands = []cli.Command{
		{
			Name: "foo",
			Action: func(c *cli.Context) error {
				program := c.App.Name
				fmt.Println("command", program)
				return nil
			},
			Subcommands: []cli.Command{
				{
					Name: "bar",
					Action: func(c *cli.Context) error {
						program := c.App.Name
						fmt.Println("subcommand", program)
						return nil
					},
				},
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		program := c.App.Name
		fmt.Println(program)
		return nil
	}

	log.Fatal(app.Run(os.Args))
}
