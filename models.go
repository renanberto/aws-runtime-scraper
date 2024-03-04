package main

type Configuration struct {
	General struct {
		Scraper ScraperConfig `yaml:"scraper"`
		AWS     AWSConfig     `yaml:"aws"`
	} `yaml:"general"`
}

type ScraperConfig struct {
	TableSelector string `yaml:"tableSelector"`
	URL           string `yaml:"url"`
	TableTitle    string `yaml:"tableTitle"`
	FilePath      string `yaml:"filePath"`
}

type AWSConfig struct {
	Accounts []AWSAccount `yaml:"accounts"`
}

type AWSAccount struct {
	ID      string `yaml:"id"`
	Token   string `yaml:"token"`
	Key     string `yaml:"key"`
	ARN     string `yaml:"arn"`
	Session string `yaml:"session"`
	MFA     string `yaml:"mfa"`
}

type TableEOL struct {
	Name                string
	Identifier          string
	OperatingSystem     string
	DeprecationDate     string
	BlockFunctionCreate string
	BlockFunctionUpdate string
}

type LambdaProperties struct {
	FunctionName string
	FunctionARN  string
	Runtime      string
	Version      string
	LastModified string
}

func convertRuntimeToArray(tableEOL []TableEOL) []string {
	runtime := []string{}
	for _, v := range tableEOL {
		runtime = append(runtime, v.Identifier)
	}
	return runtime
}
