package main

import (
	"os"
	"testing"
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
		content := loadFile("./test_resources/basic.lox")
		scanner := NewScanner(content)
		tokens := scanner.ScanTokens()

		expectedTokens := []TokenType{
			LEFT_PAREN,
			LEFT_PAREN,
			RIGHT_PAREN,
			RIGHT_PAREN,
			LEFT_BRACE,
			RIGHT_BRACE,
			BANG,
			STAR,
			PLUS,
			MINUS,
			SLASH,
			EQUAL,
			LESS,
			GREATER,
			LESS_EQUAL,
			EQUAL_EQUAL,
			EOF,
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
		content := loadFile("./test_resources/literals.lox")
		scanner := NewScanner(content)
		tokens := scanner.ScanTokens()

		expectedTokens := []TokenType{
			VAR,
			IDENTIFIER,
			EQUAL,
			STRING,
			SEMICOLON,
			FOR,
			LEFT_PAREN,
			VAR,
			IDENTIFIER,
			EQUAL,
			NUMBER,
			SEMICOLON,
			IDENTIFIER,
			LESS,
			NUMBER,
			SEMICOLON,
			IDENTIFIER,
			EQUAL,
			IDENTIFIER,
			PLUS,
			NUMBER,
			RIGHT_PAREN,
			LEFT_BRACE,
			PRINT,
			IDENTIFIER,
			SEMICOLON,
			RIGHT_BRACE,
			EOF,
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
