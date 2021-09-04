package suppliers

type Supplier struct {
	id                  int
	name                string
	email               string
	phone               string
	inn                 int
	kpp                 int
	description         string
	reputation          int
	soldAmount          float64
	successfulTenders   int
	unsuccessfulTenders int
	isInnovate          bool
}
