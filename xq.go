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

type RssItem struct {
	Title  string
	Link   string
}

type RssItems []RssItem

func (b RssItems) Len() int {
	return len(b)
}

func main() {
	app := App()
	app.Action = Action
	app.Commands = []cli.Command{
		{
			Name:    "item",
			Aliases: []string{"i"},
			Usage:   "item a list",
			Action:  func(c *cli.Context) error {
				file, _ := os.Open(c.Args().First())
				defer file.Close()
				fp := gofeed.NewParser()
				feed, _ := fp.Parse(file)
				items := feed.Items
				//fmt.Printf("%s", items)
				var RssItems RssItems
				for _, item := range items {
					var RssItem RssItem = RssItem{item.Title, item.Link}
					RssItems = append(RssItems, RssItem)
				}
				for _, item := range RssItems {
					fmt.Printf("%s : %s\n", item.Title, item.Link)
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
