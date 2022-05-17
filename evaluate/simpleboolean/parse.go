package simpleboolean

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// ParseExpression takes in an expression and parses it.
func ParseExpression(expr string) (parseCtx IParseContext, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)

			if !ok {
				err = fmt.Errorf("DSL compilation panic with non-error type %T: %v", r, r)
			}
		}
	}()

	lexer := NewSimpleBooleanLexer(antlr.NewInputStream(expr))
	lexerErrorListener := NewSimpleErrorListener(expr)
	lexer.AddErrorListener(lexerErrorListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parserErrorListener := NewSimpleErrorListener(expr)
	parser := NewSimpleBooleanParser(stream)
	parser.AddErrorListener(parserErrorListener)
	parseCtx = parser.Parse()

	if len(lexerErrorListener.Errors) > 0 {
		return nil, lexerErrorListener.Errors[0]
	}

	if len(parserErrorListener.Errors) > 0 {
		return nil, parserErrorListener.Errors[0]
	}

	return parseCtx, nil
}
