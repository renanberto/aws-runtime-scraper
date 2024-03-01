package main

type eol struct {
	table  TableEOL
	lambda LambdaProperties
}

func NewEOL(table TableEOL, lambda LambdaProperties) eol {
	return eol{
		table:  table,
		lambda: lambda,
	}
}

func (e eol) ValidateEOLByIntentifier(identifier string) {
}
