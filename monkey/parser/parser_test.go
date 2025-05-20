package parser

import (
	"testing"

	"github.com/r-priyanshu/interpreter/ast"
	"github.com/r-priyanshu/interpreter/lexer"
)

func TestLetStatement(t *testing.T) {
	input :=
		`
		let x = 3;
		let y = 10;
		let foobar = 838383; 
	`
	//get our tokens
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Does not contain 3 statements:got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}

}

func testLetStatement(t *testing.T, stmt ast.Statement, identifier string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("Not a let: got=%q", stmt.TokenLiteral())
		return false
	}
	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("Statement si not *ast.LetStatement: got=%T", stmt)
		return false
	}
	if letStmt.Name.Value != identifier {
		t.Errorf("statement value not '%s': got=%s", identifier, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != identifier {
		t.Errorf("statement value not '%s': got=%s", identifier, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("Parser has %d  errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error %q", msg)
	}
	t.FailNow()
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 45;
	return a;
	return 60347;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain three statements, got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {

		 returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}
