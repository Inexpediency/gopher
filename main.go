package main

import "github.com/ythosa/gobih/counter"

func main() {
	env := counter.Env{
		"x": 10,
		"y": 100,
	}

	counter.Count("x + y + sqrt(y) + sqrt(x) + pow(99, 99)", env)
}
