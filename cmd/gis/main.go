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
		query     string
		directory string
		filename  string
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
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "directory, d",
					Usage:       "download directory",
					Value:       "",
					Destination: &directory,
				},
				cli.StringFlag{
					Name:        "filename, f",
					Usage:       "download filename (ex: <filename>51.jpg)",
					Value:       "image",
					Destination: &filename,
				},
			},
			Action: func(c *cli.Context) error {
				if directory == "" {
					directory, _ = os.Getwd()
				}
				if f, err := os.Stat(directory); os.IsNotExist(err) || !f.IsDir() {
					fmt.Fprintf(c.App.Writer, "%s is not exists, or is not directory", directory)
					return nil
				}
				images, err := gis.Search(query)
				if err != nil {
					fmt.Fprintf(c.App.Writer, "%s\n", err.Error())
					return nil
				}
				results := gis.Download(directory, filename, images)
				for _, result := range results {
					if result.Error != nil {
						fmt.Fprintf(c.App.Writer, "%s\n", result.Error.Error())
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
