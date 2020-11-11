package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type attempt struct {
	right  string
	writed string
}

type words map[string]attempt

func parseWords(inputFilePath string) words {
	inf, err := os.OpenFile(inputFilePath, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal("Invalid input file path")
	}

	data, err := ioutil.ReadAll(inf)
	if err != nil {
		log.Fatal(err)
	}

	parsed := make(words)

	lines := strings.Split(string(data), "\n")
	for _, pair := range lines {
		p := strings.Split(pair, ": ")

		ru := strings.TrimSpace(p[0])
		en := strings.TrimSpace(p[1])

		parsed[ru] = attempt{
			right:  en,
			writed: "",
		}
	}

	return parsed
}

func memorize(w words, verbose bool) words {
	wrong := make(words)

	reader := bufio.NewReader(os.Stdin)
	for ru, a := range w {
		fmt.Printf("%s: ", ru)
		wr, _ := reader.ReadString('\n')
		wr = strings.TrimSpace(wr)
		if wr != a.right {
			if verbose {
				fmt.Printf("Must be: %s\n", a.right)
			}

			wrong[ru] = attempt{
				right:  a.right,
				writed: wr,
			}
		}
	}

	return wrong
}

func printResults(w words) {
	fmt.Println()
	defer fmt.Println()

	if len(w) == 0 {
		fmt.Println("All right!")
		return
	}

	fmt.Println("Your mistakes are:")
	for ru, a := range w {
		fmt.Printf("%s | %s | %s\n", ru, a.right, a.writed)
	}
}

func main() {
	var inputFile = flag.String("i", "./input/words.yaml", "input file path")
	var errorCorrection = flag.Bool("ec", true, "Enable error correction")
	var verbose = flag.Bool("v", true, "Verbose errors")

	words := parseWords(*inputFile)
	words = memorize(words, *verbose)
	printResults(words)

	if !*errorCorrection {
		return
	}

	for len(words) != 0 {
		words = memorize(words, *verbose)
		printResults(words)
	}
}
