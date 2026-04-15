package parser

import (
	"testing"
	"tree-walk-interpreter/token"
)

func TestAstPrinter(t *testing.T) {
	expr := NewBinary(
		NewUnary(
			token.Token{Type: token.MINUS, Lexeme: "-"},
			NewLiteral(123),
		),
		token.Token{Type: token.STAR, Lexeme: "*"},
		NewGrouping(
			NewLiteral(45.67),
		),
	)

	printer := AstPrinter{}
	result := expr.Accept(&printer)

	if result != "(* (- 123) (group 45.67))" {
		t.Errorf("expected %s, got %s", "(+ (- 123) (* 45.67))", result)
	}
}
