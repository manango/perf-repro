package simpleboolean

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

//SimpleSyntaxError represents the syntax errors raised during lexical and parsing stages.
type SimpleSyntaxError struct {
	line       int
	column     int
	msg        string
	expression string
}

// Error returns the compilation error.
func (e *SimpleSyntaxError) Error() string {
	msg := "DSL rule had an error compiling at %d:%d \n\nError: %s \n\nRule: %s"
	return fmt.Sprintf(msg, e.line, e.column, e.msg, e.expression)
}

// SimpleErrorListener used to listen to all the errors raised during lexical and parsing stages.
type SimpleErrorListener struct {
	*antlr.DefaultErrorListener
	Errors     []error
	Expression string
}

// NewSimpleErrorListener provides a new SimpleErrorListener
func NewSimpleErrorListener(expression string) *SimpleErrorListener {
	return &SimpleErrorListener{Expression: expression}
}

// SyntaxError listener method gets called on the SimpleErrorListener incase of any syntax errors.
func (l *SimpleErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException) {
	l.Errors = append(l.Errors, &SimpleSyntaxError{
		line:       line,
		column:     column,
		msg:        msg,
		expression: l.Expression,
	})
}
