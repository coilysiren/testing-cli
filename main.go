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
	app.Action = func(c *cli.Context) error {
		fmt.Println("c.App.Name for app.Action is", c.App.Name)
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name: "foo",
			Action: func(c *cli.Context) error {
				fmt.Println("c.App.Name for app.Commands.Action is", c.App.Name)
				return nil
			},
			Subcommands: []cli.Command{
				{
					Name: "bar",
					Action: func(c *cli.Context) error {
						fmt.Println("c.App.Name for App.Commands.Subcommands.Action is", c.App.Name)
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal()
	}
}
