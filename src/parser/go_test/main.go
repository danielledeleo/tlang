package main

import (
	"./parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"io/ioutil"
	"log"
	"os"
)

type turingListener struct {
	*parser.BaseTuringListener
}

func (l *turingListener) ExitVariableDeclaration(c *parser.VariableDeclarationContext) {
	identifiers := c.VariableIdentifierList()
	for _, id := range identifiers.GetChildren() {
		payload := id.GetPayload()
		fmt.Println(payload.(*parser.VariableIdentifierContext).IDENTIFIER().GetText())

	}

	//fmt.Println(identifiers, "->", c.TypeSpec().GetText())
}

func (l *turingListener) ExitExpressionContext(c *parser.ExpressionContext) {

}

func main() {
	filename := os.Args[1]
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	in := antlr.NewInputStream(string(file))

	lexer := parser.NewTuringLexer(in)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewTuringParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(&turingListener{}, p.Program())
}
