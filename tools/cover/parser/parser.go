package parser

type Parser struct {
}

func NewParser() {

}

type ParserOpts func(*Parser) error

// WithIgnoreGenerated removes generated files from the coverage report.
// detection is based on https://tinyurl.com/4j379zew
func WithIgnoreGenerated() ParserOpts {
	return func(p *Parser) error {
		p
	}
}
