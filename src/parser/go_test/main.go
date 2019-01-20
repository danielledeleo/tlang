package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"./parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type turingListener struct {
	*parser.BaseTuringListener
}

func (l *turingListener) ExitVariableDeclaration(c *parser.VariableDeclarationContext) {
	identifiers := c.VariableIdentifierList()

	for _, id := range identifiers.GetChildren() {
		fmt.Println(reflect.TypeOf(id))
	}

	fmt.Println(identifiers, "->", c.TypeSpec().GetText())
}

func (l *turingListener) ExitMultiplicativeExpression(c *parser.MultiplicativeExpressionContext) {
	if c.MultiplicativeOperator() != nil {
		fmt.Println(c.MultiplicativeOperator().GetText())
	}
}

func main() {
	stdin, err  := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	in := antlr.NewInputStream(string(stdin))

	lexer := parser.NewTuringLexer(in)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewTuringParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(&turingListener{}, p.Program())
}
