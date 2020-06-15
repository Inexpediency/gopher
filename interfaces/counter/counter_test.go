package counter_test

import (
	"fmt"
	"github.com/ythosa/gobih/interfaces/counter"
	"math"
	"testing"
)

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
	}

	var prevExpr string
	for _, test := range tests {
		// Return expr, only where it is changed
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}

		expr, err := counter.Parse(test.expr)
		if err != nil {
			t.Error(err) // Analyze error
			continue
		}
			got := fmt.Sprintf("%.6g", expr.Eval(test.env))
			fmt.Printf("\t%v => %s\n", test.env, got)
			if got != test.want{
				t.Errorf("%s.Eval() в %v = %q, must be %q\n",
					test.expr, test.env, got, test.want)
		}
	}
}
