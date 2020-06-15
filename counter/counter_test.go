package counter_test

import (
	"fmt"
	"github.com/ythosa/gobih/counter"
	"math"
	"testing"
)

//!+Eval
func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  counter.Env
		want string
	}{
		{"sqrt(A / pi)", counter.Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", counter.Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", counter.Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", counter.Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", counter.Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", counter.Env{"F": 212}, "100"},
		//!-Eval
		// additional tests that don't appear in the book
		{"-1 + -x", counter.Env{"x": 1}, "-2"},
		{"-1 - x", counter.Env{"x": 1}, "-2"},
		//!+Eval
	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := counter.Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}

//!-Eval

/*
//!+output
sqrt(A / pi)
	map[A:87616 pi:3.141592653589793] => 167
pow(x, 3) + pow(y, 3)
	map[x:12 y:1] => 1729
	map[x:9 y:10] => 1729
5 / 9 * (F - 32)
	map[F:-40] => -40
	map[F:32] => 0
	map[F:212] => 100
//!-output
// Additional outputs that don't appear in the book.
-1 - x
	map[x:1] => -2
-1 + -x
	map[x:1] => -2
*/

func TestErrors(t *testing.T) {
	for _, test := range []struct{ expr, wantErr string }{
		{"x % 2", "unexpected '%'"},
		{"math.Pi", "unexpected '.'"},
		{"!true", "unexpected '!'"},
		{`"hello"`, "unexpected '\"'"},
		{"log(10)", `unknown function "log"`},
		{"sqrt(1, 2)", "call to sqrt has 2 args, want 1"},
	} {
		expr, err := counter.Parse(test.expr)
		if err == nil {
			vars := make(map[counter.Var]bool)
			err = expr.Check(vars)
			if err == nil {
				t.Errorf("unexpected success: %s", test.expr)
				continue
			}
		}
		fmt.Printf("%-20s%v\n", test.expr, err) // (for book)
		if err.Error() != test.wantErr {
			t.Errorf("got error %s, want %s", err, test.wantErr)
		}
	}
}

/*
//!+errors
x % 2               unexpected '%'
math.Pi             unexpected '.'
!true               unexpected '!'
"hello"             unexpected '"'
log(10)             unknown function "log"
sqrt(1, 2)          call to sqrt has 2 args, want 1
//!-errors
*/
