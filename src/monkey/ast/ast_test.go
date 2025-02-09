package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	expectStringEq(t, "program.String()", "let myVar = anotherVar;", program.String())
}

func expectStringEq(t *testing.T, description string, expected string, actual string) {
	if expected != actual {
		t.Errorf("%s wrong. expected='%s' got='%s'", description, expected, actual)
	}
}
