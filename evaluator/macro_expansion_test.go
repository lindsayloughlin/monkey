package evaluator

import (
	"testing"
	"interpreter/object"
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/parser"
)

func TestDefineMacros(t *testing.T) {
	input := `

let number = 1;
let function = fn(x,y) { x + y;}
let mymacro = macro(x, y){x +y;}
`

	env := object.NewEnvironment()
	program := testParseProgram(input)
	DefineMacros(program, env)

	if len(program.Statements) != 2 {
		t.Fatalf("Wrong number of statements, got=%d", len(program.Statements))
	}

	_, ok := env.Get("number")
	if ok {
		t.Fatalf("number should not be defined")
	}
	_, ok = env.Get("function" )
	if ok {
		t.Fatalf("function should not be defined")
	}
	obj, ok := env.Get("mymacro")
	if !ok {
		t.Fatalf("object is not Macro. got=%T (%+v)", obj, obj)
	}
	macro ,ok := obj.(*object.Macro)

	if len(macro.Parameters) != 2 {
		t.Fatalf("Wrong number of macro parameters. got=%d", len(macro.Parameters))
	}
}
func testParseProgram(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}
