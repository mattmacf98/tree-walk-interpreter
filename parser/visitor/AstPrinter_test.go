package visitor

import (
	"testing"
	"tree-walk-interpreter/parser/grammar"
	"tree-walk-interpreter/token"
)

func TestAstPrinter(t *testing.T) {
	expr := grammar.NewBinary(
		grammar.NewUnary(
			token.Token{Type: token.MINUS, Lexeme: "-"},
			grammar.NewLiteral(123),
		),
		token.Token{Type: token.STAR, Lexeme: "*"},
		grammar.NewGrouping(
			grammar.NewLiteral(45.67),
		),
	)

	printer := AstPrinter{}
	result := expr.Accept(&printer)

	if result != "(* (- 123) (group 45.67))" {
		t.Errorf("expected %s, got %s", "(+ (- 123) (* 45.67))", result)
	}
}
