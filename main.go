package main

import (
	"fmt"
	"log"
	"os"

	"github.com/artheus/swaggable/antlr/parser"
	"github.com/artheus/swaggable/mapping"
	"github.com/artheus/swaggable/swagger"
	yaml "gopkg.in/yaml.v2"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*parser.BaseSwaggableParserListener
}

var components map[string]*swagger.Component = make(map[string]*swagger.Component)
var baseElements map[string]*mapping.BaseElement = make(map[string]*mapping.BaseElement)

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterObjDeclaration(ctx *parser.ObjDeclarationContext) {
	var comp = mapping.CreateComponent(ctx)
	components[comp.Name] = comp
}

func (this *TreeShapeListener) EnterEnumDeclaration(ctx *parser.EnumDeclarationContext) {
	var enum = mapping.CreateEnum(ctx)
	components[enum.Name] = enum
}

func (this *TreeShapeListener) EnterBaseDeclaration(ctx *parser.BaseDeclarationContext) {
	var base = mapping.CreateBase(ctx)
	baseElements[base.Name] = base
}

func (s *TreeShapeListener) ExitProgram(ctx *parser.ProgramContext) {
	// Base extends Base
	for _, b := range baseElements {
		for extIndex, ext := range b.GetExtends() {
			if base, ok := baseElements[ext]; ok {
				for _, baseProp := range base.GetProperties() {
					b.AddProperty(baseProp)
				}
				b.DeleteExtends(extIndex)
			}
		}
	}

	// Component extends Base
	for i, comp := range components {
		for extIndex, ext := range comp.GetExtends() {
			if base, ok := baseElements[ext]; ok {
				for _, baseProp := range base.GetProperties() {
					comp.AddProperty(baseProp)
				}
				comp.DeleteExtends(extIndex)
			}
		}
		extComp := comp.CompileExtends()
		if extComp != nil {
			components[i] = extComp
		}
	}

	root := swagger.NewRoot()
	root.Components.AddSchemas(components)

	// CRUD paths for components
	for _, comp := range components {
		for _, prop := range comp.GetProperties() {
			if prop.Indexed {
				root.AddPath(swagger.CreateGetAndDeleteByIndexPath(comp, prop))
			}
		}
		root.AddPath(swagger.CreateUpdateAndCreatePath(comp))
	}

	d, err := yaml.Marshal(&root)
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
}
