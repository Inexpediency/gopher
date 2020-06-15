package interfaces

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
	x Expr
}

// binary represents an expression with a binary operator, such as x+y
type binary struct {
	op rune // `+` or `-` or `*` or `/`
	x, y Expr
}

// call represents a function call expression, such as sin(x)
type call struct {
	fn string // one of `pow`, `sin`, `sqrt`
	args []Expr
}

// Env - an environment that maps variable names to values
type Env map[Var]float64
