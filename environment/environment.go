package environment

import (
	"fmt"
	"tree-walk-interpreter/token"
)

type Environment struct {
	values map[string]any
}

func NewEnvironment() Environment {
	return Environment{
		values: make(map[string]any),
	}
}

func (e *Environment) Define(name string, value any) {
	e.values[name] = value
}

func (e *Environment) Assign(name string, value any) error {
	if _, ok := e.values[name]; !ok {
		return fmt.Errorf("undefined variable '%s'", name)
	}
	e.values[name] = value
	return nil
}

func (e *Environment) Get(name token.Token) (any, error) {
	value, ok := e.values[name.Lexeme]
	if !ok {
		return nil, fmt.Errorf("undefined variable '%s'", name.Lexeme)
	}
	return value, nil
}
