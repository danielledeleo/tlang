package main

import (
	"fmt"

	"./parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type turingListener struct {
	*parser.BaseTuringListener
}

func (l *turingListener) ExitVariableDeclaration(c *parser.VariableDeclarationContext) {
	fmt.Println(c.VariableIdentifier().GetText(), "->", c.TypeSpec().GetText())
}

func (l *turingListener) ExitMultiplicativeExpression(c *parser.MultiplicativeExpressionContext) {
	if c.MultiplicativeOperator() != nil {
		fmt.Println(c.MultiplicativeOperator().GetText())
	}
}

func main() {
	in := antlr.NewInputStream("var x : someRecord := 5 * (3 + 5)")

	lexer := parser.NewTuringLexer(in)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewTuringParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(&turingListener{}, p.Program())
}
