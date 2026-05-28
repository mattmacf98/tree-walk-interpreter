package environment

import (
	"fmt"
	"tree-walk-interpreter/token"
)

type Environment struct {
	values    map[string]any
	enclosing *Environment
}

func NewEnvironment(enclosing *Environment) Environment {
	return Environment{
		values:    make(map[string]any),
		enclosing: enclosing,
	}
}

func (e *Environment) Define(name string, value any) {
	e.values[name] = value
}

func (e *Environment) Assign(name string, value any) error {
	if _, ok := e.values[name]; !ok {
		if e.enclosing == nil {
			return fmt.Errorf("undefined variable '%s'", name)
		} else {
			return e.enclosing.Assign(name, value)
		}

	}
	e.values[name] = value
	return nil
}

func (e *Environment) Get(name token.Token) (any, error) {
	value, ok := e.values[name.Lexeme]
	if !ok {
		if e.enclosing == nil {
			return nil, fmt.Errorf("undefined variable '%s'", name.Lexeme)
		} else {
			return e.enclosing.Get(name)
		}

	}

	return value, nil
}
