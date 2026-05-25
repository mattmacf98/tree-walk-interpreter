package visitor

import (
	"fmt"
	"os"
	"testing"
	"tree-walk-interpreter/parser"
	"tree-walk-interpreter/parser/expression"
	"tree-walk-interpreter/scanner"
	"tree-walk-interpreter/token"
)

func loadFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func TestAstPrinterGrammar(t *testing.T) {
	expr := expression.NewBinary(
		expression.NewUnary(
			token.Token{Type: token.MINUS, Lexeme: "-"},
			expression.NewLiteral(123),
		),
		token.Token{Type: token.STAR, Lexeme: "*"},
		expression.NewGrouping(
			expression.NewLiteral(45.67),
		),
	)

	printer := AstPrinter{}
	result := expr.Accept(&printer)
	expected := "(* (- 123) (group 45.67))"

	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestAstPrinterParser(t *testing.T) {
	content := loadFile("./fixtures/basic.lox")
	scanner := scanner.NewScanner(content)
	tokens := scanner.ScanTokens()

	parser := parser.NewParser(tokens)
	expr, err := parser.Parse()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	fmt.Println(expr)

	// expected := "(* (- 123) (group 45.67))"

	// printer := AstPrinter{}
	// result := expr.Accept(&printer)
	// if result != expected {
	// 	t.Errorf("expected %s, got %s", expected, result)
	// }
}
