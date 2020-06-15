package interfaces

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

/* XML decoding based on tokens
	package xml code fragment   */
/*
	type Name struct {
		Local string // For ex: `Title` or `id`
	}
	type Attr struct { // For ex: name=value
		Name Name
		Value string
	}
	// Token includes StartElement, EndElement, CharData,
	// Comment ond others ... ( they are hidden )
	type Token interface {}

	type StartElement struct { Name Name; Attr []Attr } // For ex: <name>
	type EndElement struct { Name Name } // For ex: </name>
	type CharData []byte // For ex: <name>Char data and char data</name>
	type Comment []byte // For ex: <!-- Comment -->

	type Decoder struct {}
	func NewDecoder (io.Reader) *Decoder {...}
	func (*Decoder) Token() (Token, error) // Return token
*/

func DecodeXML() {
	// Build example: curl  http://www.w3.org/TR/2006/RECxmlll20060816 | main div div h2
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // Stack of elements names
	for {
		tok, err := dec.Token()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		switch t := tok.(type) {
		case xml.StartElement:
			stack = append(stack, t.Name.Local) // Push to stack
		case xml.EndElement:
			stack = stack[:len(stack) - 1] // Remove from stack
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

// ./webworkers/fetch.go http://www.w3.org/TR/2006/RECxmlll20060816 | ./xmlselect div div h2