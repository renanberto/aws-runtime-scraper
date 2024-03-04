# AWS runtime EOL page scraper

Simple scraper to get all data form table and check ur lambdas

### TO DO

Rename `config/example.yml` to `config/general.yml`

--- 
#### Usage

Firstly should clone the repo local or fork and ajust your properties.

* `edir your properties config/general.yml`
* `go build`
* `chmod +x automatic-runtime-validate && mv automatic-runtime-validate /usr/local/bin/`
* `automatic-runtime-validate --help`

Global
```
NAME:
   AWS Lambda runtime EOL - A new cli application

USAGE:
   AWS Lambda runtime EOL [global options] command [command options] 

COMMANDS:
   search-by-region  search all lambdas EOL by region
   help, h           Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
``` 

Per Region
```NAME:
   AWS Lambda runtime EOL search-by-region - search all lambdas EOL by region

USAGE:
   AWS Lambda runtime EOL search-by-region [command options] [arguments...]

OPTIONS:
   --region value  Region to search
   --export        Export result to CSV file (default: false)
   --help, -h      show help
```