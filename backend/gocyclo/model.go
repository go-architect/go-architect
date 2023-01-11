package gocyclo

type GoCycloOutput struct {
	AverageComplexity float64
	ComplexityDetails []*ComplexityDetail
}

type ComplexityDetail struct {
	Package    string
	File       string
	Function   string
	Complexity int
}
