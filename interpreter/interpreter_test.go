package interpreter

import (
	"math"
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
		got := interp.evaluate(expr)

		v, ok := got.(float64)
		if !ok {
			t.Fatalf("expected float64, got %T (%v)", got, got)
		}

		want := -123.0 * 45.67
		if math.Abs(v-want) > 1e-9 {
			t.Errorf("expected %v, got %v", want, v)
		}
	})
}
