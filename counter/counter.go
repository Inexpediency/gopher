package counter

import (
	"fmt"
	"math"
	"strings"
)

// Expr represents an arithmetic expression
type Expr interface {
	// Eval returns the value of this Expr in the `env` environment
	Eval(env Env) float64
	// Check reports errors in this Expr and adds its own Vars
	Check(vars map[Var]bool) error
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


func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (literal) Check(vars map[Var]bool) error {
	return nil
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("invalid unary operator %q", u.op)
	}
	return u.x.Check(vars)
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("invalid binary operator %q", b.op)
	}

	if err := b.x.Check(vars); err != nil {
		return err
	}

	return b.y.Check(vars)
}

var numParams = map[string]int{
	"pow":  2,
	"sin":  1,
	"sqrt": 1,
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}

	if len(c.args) != arity {
		return fmt.Errorf("call %s has %d instead of %d", c.fn, len(c.args), arity)
	}

	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}

	return nil
}

func parseAndCheck(s string) (Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("unknown variable: %s", v)
		}
	}
	return expr, nil
}


func Count(s string, env Env) {
	expr, err := parseAndCheck(s)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(expr.Eval(env))
}
