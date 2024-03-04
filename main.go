package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		EnableBashCompletion: true,
		Name:                 "AWS Lambda runtime EOL",
		Commands: []*cli.Command{
			{
				Name:  "search-by-region",
				Usage: "search all lambdas EOL by region",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "region",
						Usage:    "Region to search",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "config-file",
						Usage:    "Load configurion file",
						Required: true,
					},
					&cli.BoolFlag{
						Name:  "export",
						Usage: "Export result to CSV file",
					},
				},
				Action: func(cCtx *cli.Context) error {
					configure(cCtx.String("config-file"))

					scrapper := NewScraper(config).Run()
					awss := NewAWS(config, scrapper)

					awss.SearchRuntimeByRegion(cCtx.String("region"))
					if cCtx.Bool("export") {
						scrapper.toCSV()
					}
					return nil
				},
			},
			{
				Name:  "search-all",
				Usage: "search all lambdas EOL at all regions (low performance)",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "config-file",
						Usage:    "Load configurion file",
						Required: true,
					},
				},
				Action: func(cCtx *cli.Context) error {
					configure(cCtx.String("config-file"))

					scrapper := NewScraper(config).Run()
					awss := NewAWS(config, scrapper)

					awss.SearchRuntimeAllRegions()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
