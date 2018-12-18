package main

import (
	"os"

	"github.com/o-sk/gis"
	"github.com/urfave/cli"
)

func display(images []gis.Image) {

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

	app.Action = func(conrext *cli.Context) error {
		images := gis.Search(query)
		display(images)
		return nil
	}
	app.Run(os.Args)
}