package lexemcounter

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type visitor struct {
	Package map[*ast.GenDecl]bool
}

var gcount map[int][]string
var lcount map[int][]string

func unique(s []string) []string {
	keys := make(map[string]bool)
	list := make([]string, 0)
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func addGCount(v string) {
	if gcount == nil {
		gcount = make(map[int][]string)
	}
	if gcount[len(v)] == nil {
		gcount[len(v)] = []string{v}
	} else {
		gcount[len(v)] = append(gcount[len(v)], v)
	}
}

func addLCount(v string) {
	if lcount == nil {
		lcount = make(map[int][]string)
	}
	if lcount[len(v)] == nil {
		lcount[len(v)] = []string{v}
	} else {
		lcount[len(v)] = append(lcount[len(v)], v)
	}
}

func printGCount() {
	fmt.Println("GLOBAL VARIABLES:")

	for k, v := range gcount {
		fmt.Printf("\tLength: %d \tCount: %d \tVariables:", k, len(v))
		for _, i := range unique(v) {
			fmt.Printf(" %s", i)
		}
		fmt.Println()
	}
}

func printLCount() {
	fmt.Println("LOCAL VARIABLES:")

	for k, v := range lcount {
		fmt.Printf("\tLength: %d \tCount: %d \tVariables:", k, len(v))
		for _, i := range unique(v) {
			fmt.Printf(" %s", i)
		}
		fmt.Println()
	}
}

func makeVisitor(f *ast.File) visitor {
	k1 := make(map[*ast.GenDecl]bool)
	for _, aa := range f.Decls {
		v, ok := aa.(*ast.GenDecl)
		if ok {
			k1[v] = true
		}
	}

	return visitor{Package: k1}
}

func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	switch d := n.(type) {
	case *ast.AssignStmt:
		if d.Tok != token.DEFINE {
			return v
		}

		for _, name := range d.Lhs {
			v.isItLocal(name)
		}
	case *ast.RangeStmt:
		v.isItLocal(d.Key)
		v.isItLocal(d.Value)
	case *ast.FuncDecl:
		if d.Recv != nil {
			v.checkAll(d.Recv.List)
		}
		v.checkAll(d.Type.Params.List)
		if d.Type.Results != nil {
			v.checkAll(d.Type.Results.List)
		}
	case *ast.GenDecl:
		if d.Tok != token.VAR {
			return v
		}
		for _, spec := range d.Specs {
			value, ok := spec.(*ast.ValueSpec)
			if ok {
				for _, name := range value.Names {
					if name.Name == "_" {
						continue
					}
					if v.Package[d] {
						addGCount(name.Name)
					} else {
						addLCount(name.Name)
					}
				}
			}
		}
	}

	return v
}

func (v visitor) isItLocal(n ast.Node) {
	identifier, ok := n.(*ast.Ident)
	if !ok {
		return
	}

	if identifier.Name == "_" || identifier.Name == "" {
		return
	}

	if identifier.Obj != nil && identifier.Obj.Pos() == identifier.Pos() {
		addLCount(identifier.Name)
	}
}

func (v visitor) checkAll(fs []*ast.Field) {
	for _, f := range fs {
		for _, name := range f.Names {
			v.isItLocal(name)
		}
	}
}

func CountVariablesOfLen(paths []string) {
	var v visitor
	all := token.NewFileSet()
	for _, file := range paths {
		fmt.Println("Processing file: ", file)
		f, err := parser.ParseFile(all, file, nil, parser.AllErrors)
		if err != nil {
			fmt.Println(err)
			continue
		}
		v = makeVisitor(f)
		ast.Walk(v, f)

		printGCount()
		printLCount()

		gcount = nil
		lcount = nil
	}
}
