package counter

import (
	"fmt"
	"math"
)

// Expr represents an arithmetic expression
type Expr interface {
	// Eval returns the value of this Expr in the `env` environment
	Eval(env Env) float64
}

// Var defines a variable, such as x
type Var string

// literal is a numeric constant, such as 3.141
type literal float64

// unary represents an expression with a unary operator, such as -x
type unary struct {
	op rune // `+` or `-`
	x  Expr
}

// binary represents an expression with a binary operator, such as x+y
type binary struct {
	op   rune // `+` or `-` or `*` or `/`
	x, y Expr
}

// call represents a function call expression, such as sin(x)
type call struct {
	fn string // one of `pow`, `sin`, `sqrt`
	args []Expr
}

// Env - an environment that maps variable names to values
type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval (_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}


