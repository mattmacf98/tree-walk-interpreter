package interpreter

import (
	"fmt"
	"tree-walk-interpreter/environment"
	"tree-walk-interpreter/lox"
	"tree-walk-interpreter/parser/expression"
	"tree-walk-interpreter/parser/statement"
	"tree-walk-interpreter/token"
)

type Interpreter struct {
	environment environment.Environment
}

func NewInterpreter() Interpreter {
	return Interpreter{
		environment: environment.NewEnvironment(nil),
	}
}

func (i *Interpreter) Interpret(statements []statement.Stmt) error {
	for _, stmt := range statements {
		err := i.execute(stmt)
		if err != nil {
			lox.Error(0, err.Error())
			return err
		}
	}
	return nil
}

func (i *Interpreter) execute(stmt statement.Stmt) error {
	result := stmt.Accept(i)
	switch result := result.(type) {
	case nil:
		return nil
	case error:
		return result
	default:
		fmt.Println(i.stringify(result))
		return nil
	}
}

func (i *Interpreter) VisitExpressionStmt(stmt statement.ExpressionStmt) any {
	i.evaluate(stmt.Expression)
	return nil
}

func (i *Interpreter) VisitPrintStmt(stmt statement.PrintStmt) any {
	value := i.evaluate(stmt.Expression)
	fmt.Println(i.stringify(value))
	return nil
}

func (i *Interpreter) VisitVarStmt(stmt statement.VarStmt) any {
	value := any(nil)
	if stmt.Right != nil {
		value = i.evaluate(stmt.Right)
	}
	i.environment.Define(stmt.Name.Lexeme, value)
	return nil
}

func (i *Interpreter) VisitBlockStmt(stmt statement.BlockStmt) any {
	tempParent := i.environment
	blockEnvironment := environment.NewEnvironment(&tempParent)
	err := i.executeBlock(stmt.Statements, &blockEnvironment)
	if err != nil {
		return err
	}
	return nil
}

func (i *Interpreter) executeBlock(statements []statement.Stmt, environment *environment.Environment) error {
	previous := i.environment
	defer func() {
		i.environment = previous
	}()

	i.environment = *environment
	for _, statement := range statements {
		err := i.execute(statement)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Interpreter) VisitLiteralExpr(expr expression.Literal) any {
	return expr.Value
}

func (i *Interpreter) VisitAssignExpr(expr expression.Assign) any {
	value := i.evaluate(expr.Value)
	err := i.environment.Assign(expr.Name.Lexeme, value)
	if err != nil {
		lox.Error(expr.Name.Line, err.Error())
		return nil
	}
	return value
}

func (i *Interpreter) VisitVariableExpr(expr expression.Variable) any {
	value, err := i.environment.Get(expr.Name)
	if err != nil {
		lox.Error(expr.Name.Line, err.Error())
		return nil
	}
	return value
}

func (i *Interpreter) VisitGroupingExpr(expr expression.Grouping) any {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) VisitUnaryExpr(expr expression.Unary) any {
	right := i.evaluate(expr.Right)

	switch expr.Operator.Type {
	case token.MINUS:
		i.checkNumberOperand(expr.Operator, right)
		return -right.(float64)
	case token.BANG:
		return !i.isTruthy(right)
	}

	// should never happen
	lox.Error(expr.Operator.Line, "Unexpected unary operator.")
	return nil
}

func (i *Interpreter) VisitBinaryExpr(expr expression.Binary) any {
	left := i.evaluate(expr.Left)
	right := i.evaluate(expr.Right)

	switch expr.Operator.Type {
	case token.PLUS:
		if leftString, ok := left.(string); ok {
			if rightString, ok := right.(string); ok {
				return leftString + rightString
			}
		}
		if leftFloat, ok := left.(float64); ok {
			if rightFloat, ok := right.(float64); ok {
				return leftFloat + rightFloat
			}
		}
		lox.Error(expr.Operator.Line, "Operands must be numbers.")
		return nil
	case token.MINUS:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) - right.(float64)
	case token.STAR:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) * right.(float64)
	case token.SLASH:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) / right.(float64)
	case token.GREATER:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) > right.(float64)
	case token.GREATER_EQUAL:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) >= right.(float64)
	case token.LESS:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) < right.(float64)
	case token.LESS_EQUAL:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) <= right.(float64)
	case token.BANG_EQUAL:
		return !i.isEqual(left, right)
	case token.EQUAL_EQUAL:
		return i.isEqual(left, right)
	}

	lox.Error(expr.Operator.Line, "Unexpected binary operator.")
	return nil
}

func (i *Interpreter) evaluate(expr expression.Expr) any {
	return expr.Accept(i)
}

func (i *Interpreter) checkNumberOperand(operator token.Token, operand any) {
	_, ok := operand.(float64)
	if !ok {
		lox.Error(operator.Line, "Operand must be a number.")
	}
}

func (i *Interpreter) checkNumberOperands(operator token.Token, left, right any) {
	i.checkNumberOperand(operator, left)
	i.checkNumberOperand(operator, right)
}

func (i *Interpreter) isTruthy(value any) bool {
	if value == nil {
		return false
	}
	if value.(bool) == false {
		return false
	}
	return true
}

func (i *Interpreter) isEqual(a, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return false
	}
	return a == b
}

func (i *Interpreter) stringify(value any) string {
	if value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", value)
}
