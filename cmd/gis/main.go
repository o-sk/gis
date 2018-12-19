package main

import (
	"fmt"
	"os"

	"github.com/k0kubun/pp"
	"github.com/o-sk/gis"
	"github.com/urfave/cli"
)

func display(images []gis.Image) {
	pp.Print(images)
}

func main() {
	app := cli.NewApp()

	app.Name = "gis"
	app.Usage = "Google Image Search"
	app.Version = "0.0.1"

	var (
		query string
	)

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "query, q",
			Usage:       "Query",
			Value:       "apple",
			Destination: &query,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "download",
			Aliases: []string{"d"},
			Usage:   "download image file",
			Action: func(c *cli.Context) error {
				images, err := gis.Search(query)
				if err != nil {
					fmt.Errorf(err.Error())
					return nil
				}
				results := gis.Download(images)
				for _, result := range results {
					if result.Error != nil {
						fmt.Errorf(result.Error.Error())
					}
				}
				return nil
			},
		},
	}

	app.Action = func(conrext *cli.Context) error {
		images, err := gis.Search(query)
		if err != nil {
			fmt.Errorf(err.Error())
			return nil
		}
		display(images)
		return nil
	}
	app.Run(os.Args)
}
