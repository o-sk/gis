package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/o-sk/gis"
	"github.com/urfave/cli"
)

type json_image struct {
	Source      string
	Description string
	Cite        string
}

func display(images []gis.Image) error {
	is := make([]json_image, len(images))
	for i, image := range images {
		is[i] = json_image{
			Source:      image.Source,
			Description: image.Description,
			Cite:        image.Cite,
		}
	}
	j, err := json.Marshal(is)
	if err != nil {
		return err
	}
	fmt.Printf("%s", j)
	return nil
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

	app.Action = func(context *cli.Context) error {
		images, err := gis.Search(query)
		if err != nil {
			fmt.Fprintf(context.App.Writer, err.Error())
			return nil
		}
		err = display(images)
		if err != nil {
			fmt.Fprintf(context.App.Writer, err.Error())
			return nil
		}
		return nil
	}
	app.Run(os.Args)
}
