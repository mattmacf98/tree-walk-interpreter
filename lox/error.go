package lox

import (
	"fmt"
	"tree-walk-interpreter/token"
)

var HadError = false

func Error(line int, message string) {
	report(line, "", message)
}

func ErrorWithToken(tok token.Token, message string) {
	if tok.Type == token.EOF {
		report(tok.Line, "at end ", message)
	} else {
		report(tok.Line, fmt.Sprintf(" at '%s'", tok.Lexeme), message)
	}
}

func report(line int, where string, message string) {
	HadError = true
	fmt.Printf("[line %d] Error%s: %s\n", line, where, message)
}
