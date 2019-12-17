package main

import (
	"fmt"
	"os"
	"encoding/json"
	"github.com/urfave/cli/v2"
	_ "reflect"
	gofeed "github.com/mmcdole/gofeed"
)

func App() *cli.App {
	app := cli.NewApp()
	app.Name = "xq"
	app.Usage = "xq /path/to/rss.xml"
	app.Version = "0.2.0"
	return app
}

type RssItem struct {
	Title		string `json:"title"`
	Link		string `json:"link"`
	Updated		string `json:"update"`
	Published	string `json:"publish"`
}

type RssItems []RssItem

func (b RssItems) Len() int {
	return len(b)
}

func Action(c *cli.Context) {
	app := App()
	if c.Args().Get(0) == "" {
		help := []string{"", "--help"}
		app.Run(help)
		os.Exit(1)
	} else {
		file, _ := os.Open(c.Args().First())
		defer file.Close()
		fp := gofeed.NewParser()
		feed, _ := fp.Parse(file)
		items := feed.Items
		outputJson, err := json.Marshal(&items)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", string(outputJson))
	}
	return
}

func main() {
	app := App()
	app.Commands = []*cli.Command{
		{
			Name:    "item",
			Aliases: []string{"i"},
			Usage:   "xq i ./index.xml #output [title,link,update,publish]",
			Action:  func(c *cli.Context) error {
				file, _ := os.Open(c.Args().First())
				defer file.Close()
				fp := gofeed.NewParser()
				feed, _ := fp.Parse(file)
				items := feed.Items
				var RssItems RssItems
				for _, item := range items {
					var RssItem RssItem = RssItem{item.Title, item.Link, item.Updated, item.Published}
					RssItems = append(RssItems, RssItem)
				}
				for _, item := range RssItems {
					fmt.Printf("{\"title\":\"%s\",\"link\":\"%s\",\"update\":\"%s\",\"publish\":\"%s\"}\n", item.Title, item.Link, item.Updated, item.Published)
				}

				return nil
			},
		},
		{
			Name:    "latest",
			Aliases: []string{"l"},
			Usage:   "xq l ./index.xml #latest updated",
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
