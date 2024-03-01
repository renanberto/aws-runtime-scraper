package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	scrapper := NewScraper(config).Run()
	table := scrapper.table

	// awss := NewAWS(config)

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
					&cli.BoolFlag{
						Name:  "export",
						Usage: "Export result to CSV file",
					},
				},
				Action: func(cCtx *cli.Context) error {
					for _, v := range table {
						fmt.Println(v.Identifier)
					}
					// awss.SearchRuntimeByRegion(cCtx.String("region"))
					// if cCtx.Bool("export") {
					// 	scrapper.toCSV()
					// }
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
