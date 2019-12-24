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
    return app
}

type RssItem struct {
    Title	string `json:"title"`
    Link	string `json:"link"`
    Updated	string `json:"update"`
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
    }
    return
}

func main() {
    app := &cli.App{
	Version: "0.3.3",
	Name: "xq",
	Usage: "$ xq index.xml #xml -> json",
	Action: func(c *cli.Context) error {
	    if c.Args().Get(0) == "" {
		help := []string{"xq", "--help"}
		fmt.Printf("%s", help)
	    } else {
		file, _ := os.Open(c.Args().Get(0))
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
	    return nil
	},
    }
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
		for i, item := range RssItems {
		    if i == 0 {
			fmt.Printf("[")
		    }
		    if i == len(items) - 1 {
			fmt.Printf("{\"title\":\"%s\",\"link\":\"%s\",\"update\":\"%s\",\"publish\":\"%s\"}]", item.Title, item.Link, item.Updated, item.Published)
		    } else {
			fmt.Printf("{\"title\":\"%s\",\"link\":\"%s\",\"update\":\"%s\",\"publish\":\"%s\"},", item.Title, item.Link, item.Updated, item.Published)
		    }
		}
		return nil
	    },
	},
	{
	    Name:    "update",
	    Aliases: []string{"u","l"},
	    Usage:   "xq u ./index.xml #latest updated",
	    Subcommands: []*cli.Command{
		{
		    Name:  "link",
		    Aliases: []string{"l"},
		    Usage: "xq l link ./index.xml #latest item link",
		    Action: func(c *cli.Context) error {
			file, _ := os.Open(c.Args().First())
			defer file.Close()
			fp := gofeed.NewParser()
			feed, _ := fp.Parse(file)
			item := feed.Items[0].Link
			fmt.Printf("%s\n",item)
			return nil
		    },
		},
		{
		    Name:  "title",
		    Aliases: []string{"t"},
		    Usage: "xq l title ./index.xml #latest itme title",
		    Action: func(c *cli.Context) error {
			file, _ := os.Open(c.Args().First())
			defer file.Close()
			fp := gofeed.NewParser()
			feed, _ := fp.Parse(file)
			item := feed.Items[0].Title
			fmt.Printf("%s\n",item)
			return nil
		    },
		},
		{
		    Name:  "description",
		    Aliases: []string{"d"},
		    Usage: "xq l description ./index.xml #latest itme description",
		    Action: func(c *cli.Context) error {
			file, _ := os.Open(c.Args().First())
			defer file.Close()
			fp := gofeed.NewParser()
			feed, _ := fp.Parse(file)
			item := feed.Items[0].Description
			fmt.Printf("%s\n",item)
			return nil
		    },
		},
		{
		    Name:  "published",
		    Aliases: []string{"p"},
		    Usage: "xq l published ./index.xml #latest itme published",
		    Action: func(c *cli.Context) error {
			file, _ := os.Open(c.Args().First())
			defer file.Close()
			fp := gofeed.NewParser()
			feed, _ := fp.Parse(file)
			item := feed.Items[0].Published
			fmt.Printf("%s\n",item)
			return nil
		    },
		},
	    },
	    Action:  func(c *cli.Context) error {
		file, _ := os.Open(c.Args().First())
		defer file.Close()
		fp := gofeed.NewParser()
		feed, _ := fp.Parse(file)
		fmt.Println(feed.Updated)
		return nil
	    },
	},
	{
	    Name:    "publish",
	    Aliases: []string{"p"},
	    Usage:   "xq p ./index.xml #latest items published",
	    Action:  func(c *cli.Context) error {
		file, _ := os.Open(c.Args().First())
		defer file.Close()
		fp := gofeed.NewParser()
		feed, _ := fp.Parse(file)
		item := feed.Items[0].Published
		fmt.Printf("%s\n",item)
		return nil
	    },
	},
	{
	    Name:    "all",
	    Aliases: []string{"a"},
	    Usage:   "xq a ./index.xml #src date",
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
    }
    app.Run(os.Args)
}
