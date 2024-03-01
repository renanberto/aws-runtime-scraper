package main

type EOL interface {
	ValidateEOLByIntentifier(identifier string)
}

type Scraper interface {
	Run() SScraper
}

type AWS interface {
	SearchRuntimeAllRegions()
	SearchRuntimeByRegion(region string)
}
