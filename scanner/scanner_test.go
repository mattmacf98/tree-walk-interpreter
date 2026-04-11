package scanner

import (
	"os"
	"testing"
	"tree-walk-interpreter/token"
)

func loadFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func TestScanner(t *testing.T) {
	t.Run("basic scanner", func(t *testing.T) {
		content := loadFile("./fixtures/basic.lox")
		scanner := NewScanner(content)
		tokens := scanner.ScanTokens()

		expectedTokens := []token.TokenType{
			token.LEFT_PAREN,
			token.LEFT_PAREN,
			token.RIGHT_PAREN,
			token.RIGHT_PAREN,
			token.LEFT_BRACE,
			token.RIGHT_BRACE,
			token.BANG,
			token.STAR,
			token.PLUS,
			token.MINUS,
			token.SLASH,
			token.EQUAL,
			token.LESS,
			token.GREATER,
			token.LESS_EQUAL,
			token.EQUAL_EQUAL,
			token.EOF,
		}

		if len(tokens) != len(expectedTokens) {
			t.Fatalf("expected %d tokens, got %d", len(expectedTokens), len(tokens))
		}

		for i, want := range expectedTokens {
			if tokens[i].Type != want {
				t.Errorf("token %d: want %v, got %v", i, want, tokens[i].Type)
			}
		}
	})

	t.Run("literal scanner", func(t *testing.T) {
		content := loadFile("./fixtures/literals.lox")
		scanner := NewScanner(content)
		tokens := scanner.ScanTokens()

		expectedTokens := []token.TokenType{
			token.VAR,
			token.IDENTIFIER,
			token.EQUAL,
			token.STRING,
			token.SEMICOLON,
			token.FOR,
			token.LEFT_PAREN,
			token.VAR,
			token.IDENTIFIER,
			token.EQUAL,
			token.NUMBER,
			token.SEMICOLON,
			token.IDENTIFIER,
			token.LESS,
			token.NUMBER,
			token.SEMICOLON,
			token.IDENTIFIER,
			token.EQUAL,
			token.IDENTIFIER,
			token.PLUS,
			token.NUMBER,
			token.RIGHT_PAREN,
			token.LEFT_BRACE,
			token.PRINT,
			token.IDENTIFIER,
			token.SEMICOLON,
			token.RIGHT_BRACE,
			token.EOF,
		}

		if len(tokens) != len(expectedTokens) {
			t.Fatalf("expected %d tokens, got %d", len(expectedTokens), len(tokens))
		}

		for i, want := range expectedTokens {
			if tokens[i].Type != want {
				t.Errorf("token %d: want %v, got %v", i, want, tokens[i].Type)
			}
		}

		var testIdentifier = tokens[1]
		if testIdentifier.Literal != "test" {
			t.Errorf("test identifier: want %v, got %v", "test", testIdentifier.Literal)
		}

		var stringLiteral = tokens[3]
		if stringLiteral.Literal != "hello" {
			t.Errorf("string literal: want %v, got %v", "hello", stringLiteral.Literal)
		}

	})
}
