package interpreter

import (
	"os"
	"testing"
	"tree-walk-interpreter/parser"
	"tree-walk-interpreter/scanner"
)

func loadFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func TestInterpreter(t *testing.T) {
	t.Run("basic interpreter", func(t *testing.T) {
		content := loadFile("./fixtures/basic.lox")
		scnnr := scanner.NewScanner(content)
		tokens := scnnr.ScanTokens()
		prsr := parser.NewParser(tokens)
		expr, err := prsr.Parse()
		if err != nil {
			t.Fatalf("parse: %v", err)
		}

		interp := &Interpreter{}
		err = interp.Interpret(expr)
		if err != nil {
			t.Fatalf("interpret: %v", err)
		}
	})

	t.Run("print interpreter", func(t *testing.T) {
		content := loadFile("./fixtures/print_test.lox")
		scnnr := scanner.NewScanner(content)
		tokens := scnnr.ScanTokens()
		prsr := parser.NewParser(tokens)
		statements, err := prsr.Parse()
		if err != nil {
			t.Fatalf("parse: %v", err)
		}
		interp := &Interpreter{}
		err = interp.Interpret(statements)
		if err != nil {
			t.Fatalf("interpret: %v", err)
		}
	})
}
