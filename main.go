package main

import (
	"fmt"
	"log"
	"os"

	"github.com/artheus/swaggable/antlr/parser"

	"gopkg.in/yaml.v2"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*parser.BaseSwaggableParserListener
}

var components map[string]interface{} = make(map[string]interface{})

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterObjDeclaration(ctx *parser.ObjDeclarationContext) {
	var comp = CreateComponent(ctx)
	// TODO: Compile component extentions here.
	components[comp.Name] = comp
}

func (this *TreeShapeListener) EnterEnumDeclaration(ctx *parser.EnumDeclarationContext) {
	var enum = CreateEnum(ctx)

	components[enum.Name] = enum
}

func (s *TreeShapeListener) ExitProgram(ctx *parser.ProgramContext) {
	d, err := yaml.Marshal(&components)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(string(d))
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewSwaggableLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSwaggableParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Program()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
	fmt.Println("Done")
}
