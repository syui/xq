package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
	_ "reflect"
	gofeed "github.com/mmcdole/gofeed"
)

func Action(c *cli.Context) {
	app := App()
	if c.Args().Get(0) == "" {
		help := []string{"", "--help"}
		app.Run(help)
		os.Exit(1)
	}
}

func App() *cli.App {
	app := cli.NewApp()
	app.Name = "xq"
	app.Usage = "xq title /path/to/rss.xml"
	app.Version = "0.0.2"
	app.Author = "syui"
	return app
}

func main() {
	app := App()
	app.Action = Action
	app.Commands = []cli.Command{
		{
			Name:    "title",
			Aliases: []string{"t"},
			Usage:   "title a task on the list",
			Action:  func(c *cli.Context) error {
				file, _ := os.Open(c.Args().First())
				defer file.Close()
				fp := gofeed.NewParser()
				feed, _ := fp.Parse(file)
				items := feed.Items
				fmt.Printf("%s", items)
				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action:  func(c *cli.Context) error {
				return nil
			},
		},
	}
	app.Run(os.Args)
}
