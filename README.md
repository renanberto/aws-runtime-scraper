# AWS runtime EOL page scraper

Simple scraper to get all data form table and check ur lambdas

--- 
#### Usage

Firstly should clone the repo local or fork and ajust your properties.

* `Edir your properties like config/example.yml`
* `download released binary`
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
   search-all        search all lambdas EOL at all regions (low performance)
   help, h           Shows a list of commands or help for one command       

GLOBAL OPTIONS:
   --help, -h  show help
``` 

All regions
```
NAME:
   AWS Lambda runtime EOL search-all - search all lambdas EOL at all regions (low performance)

USAGE:
   AWS Lambda runtime EOL search-all [command options] [arguments...]

OPTIONS:
   --config-file value  Load configurion file
   --help, -h           show help
```

Per region
```
NAME:
   AWS Lambda runtime EOL search-by-region - search all lambdas EOL by region

USAGE:
   AWS Lambda runtime EOL search-by-region [command options] [arguments...]

OPTIONS:
   --region value       Region to search
   --config-file value  Load configurion file
   --export             Export result to CSV file (default: false)
   --help, -h           show help
```