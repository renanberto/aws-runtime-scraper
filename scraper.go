package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/gocolly/colly/v2"
)

type SScraper struct {
	config    ScraperConfig
	collector *colly.Collector
	table     []TableEOL
}

func NewScraper(config Configuration) Scraper {
	return SScraper{
		config:    config.General.Scraper,
		collector: colly.NewCollector(),
		table:     []TableEOL{},
	}
}

func (s SScraper) Run() SScraper {
	// Must exist to convert row in the array of struct
	schema := reflect.ValueOf(&TableEOL{}).Elem()

	// Search for tableSelector and get all rows
	s.collector.OnHTML(s.config.TableSelector, func(e *colly.HTMLElement) {
		e.ForEach("table", func(_ int, tr *colly.HTMLElement) {
			if tr.ChildText("div.title") == s.config.TableTitle {
				e.ForEach("tr", func(_ int, tr *colly.HTMLElement) {
					row := []string{}
					tr.ForEach("td", func(_ int, td *colly.HTMLElement) {
						row = append(row, strings.TrimSpace(td.Text))
					})
					for i := 0; i < schema.NumField(); i++ {
						if len(row) > i {
							if len(row[i]) == 0 {
								schema.Field(i).SetString("N/D")
								continue
							}
							if schema.Field(i).Kind() == reflect.String {
								schema.Field(i).SetString(row[i])
							}
						}
					}
					if val, ok := schema.Interface().(TableEOL); ok {
						if len(val.Name) > 0 {
							s.table = append(s.table, schema.Interface().(TableEOL))
						}
					}
				})
			}
		})
	})

	if err := s.collector.Visit(s.config.URL); err != nil {
		log.Fatal(err)
	}

	return s
}

func (s SScraper) toCSV() {
	file, err := os.Create(s.config.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	schema := reflect.TypeOf(s.table).Elem()
	var fields []string
	for i := 0; i < schema.NumField(); i++ {
		fields = append(fields, schema.Field(i).Name)
	}
	if err := writer.Write(fields); err != nil {
		log.Fatal(err)
	}

	values := reflect.ValueOf(s.table)
	for i := 0; i < values.Len(); i++ {
		var row []string
		for j := 0; j < schema.NumField(); j++ {
			row = append(row, fmt.Sprintf("%v", values.Index(i).Field(j).Interface()))
		}
		if err := writer.Write(row); err != nil {
			log.Fatal(err)
		}
	}
}
