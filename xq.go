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
	app.Usage = "xq i /path/to/rss.xml"
	app.Version = "0.0.3"
	app.Author = "syui"
	return app
}

type RssItem struct {
	Title		string
	Link		string
	Updated		string
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
			Usage:   "title, link",
			Action:  func(c *cli.Context) error {
				file, _ := os.Open(c.Args().First())
				defer file.Close()
				fp := gofeed.NewParser()
				feed, _ := fp.Parse(file)
				items := feed.Items
				var RssItems RssItems
				for _, item := range items {
					var RssItem RssItem = RssItem{item.Title, item.Link, item.Updated}
					RssItems = append(RssItems, RssItem)
				}
				for _, item := range RssItems {
					fmt.Printf("{\"title\":\"%s\",\"link\":\"%s\",\"date\":\"%s\"}\n", item.Title, item.Link, item.Updated)
				}
				return nil
			},
		},
		{
			Name:    "latest",
			Aliases: []string{"l"},
			Usage:   "latest updated",
			Action:  func(c *cli.Context) error {
				file, _ := os.Open(c.Args().First())
				defer file.Close()
				fp := gofeed.NewParser()
				feed, _ := fp.Parse(file)
				fmt.Println(feed.Updated)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
