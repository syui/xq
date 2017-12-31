package main

import (
	"fmt"
	"os"
	gofeed "github.com/mmcdole/gofeed"
	_ "reflect"
	"github.com/urfave/cli"
)

func Action(c *cli.Context) {
	app := App()
	if c.Args().Get(0) == "" {
		help := []string{"", "--help"}
		app.Run(help)
		os.Exit(1)
	}
	file, _ := os.Open(c.Args().Get(0))
	defer file.Close()
	fp := gofeed.NewParser()
	feed, _ := fp.Parse(file)
	items := feed.Items
	fmt.Printf("%s", items)
}

func App() *cli.App {
	app := cli.NewApp()
	app.Name = "xq"
	app.Usage = "xq /path/to/rss.xml"
	app.Version = "0.0.1"
	app.Author = "syui"
	return app
}

func main() {
	app := App()
	app.Action = Action
	app.Run(os.Args)
}
