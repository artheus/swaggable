// Code generated from SwaggableParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // SwaggableParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 39, 168,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 5, 2, 44, 10, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 7, 6, 65, 10, 6, 12, 6, 14, 6, 68,
	11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7, 7, 75, 10, 7, 12, 7, 14, 7, 78,
	11, 7, 3, 8, 3, 8, 3, 8, 3, 9, 5, 9, 84, 10, 9, 3, 9, 3, 9, 7, 9, 88, 10,
	9, 12, 9, 14, 9, 91, 11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 5, 10, 100, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 5, 13, 116, 10, 13, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 134, 10, 15, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 17, 3, 17, 5, 17, 142, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18,
	6, 18, 148, 10, 18, 13, 18, 14, 18, 149, 3, 18, 3, 18, 5, 18, 154, 10,
	18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 5, 20, 161, 10, 20, 3, 21, 6, 21,
	164, 10, 21, 13, 21, 14, 21, 165, 3, 21, 2, 2, 22, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 2, 2, 2, 176, 2,
	43, 3, 2, 2, 2, 4, 47, 3, 2, 2, 2, 6, 51, 3, 2, 2, 2, 8, 55, 3, 2, 2, 2,
	10, 62, 3, 2, 2, 2, 12, 71, 3, 2, 2, 2, 14, 79, 3, 2, 2, 2, 16, 83, 3,
	2, 2, 2, 18, 99, 3, 2, 2, 2, 20, 101, 3, 2, 2, 2, 22, 106, 3, 2, 2, 2,
	24, 115, 3, 2, 2, 2, 26, 117, 3, 2, 2, 2, 28, 133, 3, 2, 2, 2, 30, 135,
	3, 2, 2, 2, 32, 141, 3, 2, 2, 2, 34, 143, 3, 2, 2, 2, 36, 155, 3, 2, 2,
	2, 38, 160, 3, 2, 2, 2, 40, 163, 3, 2, 2, 2, 42, 44, 5, 40, 21, 2, 43,
	42, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 46, 7, 2, 2,
	3, 46, 3, 3, 2, 2, 2, 47, 48, 7, 13, 2, 2, 48, 49, 7, 36, 2, 2, 49, 50,
	5, 16, 9, 2, 50, 5, 3, 2, 2, 2, 51, 52, 7, 14, 2, 2, 52, 53, 7, 36, 2,
	2, 53, 54, 5, 16, 9, 2, 54, 7, 3, 2, 2, 2, 55, 56, 7, 15, 2, 2, 56, 57,
	7, 36, 2, 2, 57, 58, 7, 3, 2, 2, 58, 59, 5, 24, 13, 2, 59, 60, 7, 4, 2,
	2, 60, 61, 5, 10, 6, 2, 61, 9, 3, 2, 2, 2, 62, 66, 7, 7, 2, 2, 63, 65,
	7, 36, 2, 2, 64, 63, 3, 2, 2, 2, 65, 68, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2,
	66, 67, 3, 2, 2, 2, 67, 69, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 70, 7,
	8, 2, 2, 70, 11, 3, 2, 2, 2, 71, 76, 7, 36, 2, 2, 72, 73, 7, 10, 2, 2,
	73, 75, 7, 36, 2, 2, 74, 72, 3, 2, 2, 2, 75, 78, 3, 2, 2, 2, 76, 74, 3,
	2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 13, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 79,
	80, 7, 3, 2, 2, 80, 81, 5, 12, 7, 2, 81, 15, 3, 2, 2, 2, 82, 84, 5, 14,
	8, 2, 83, 82, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 89,
	7, 7, 2, 2, 86, 88, 5, 36, 19, 2, 87, 86, 3, 2, 2, 2, 88, 91, 3, 2, 2,
	2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 92, 3, 2, 2, 2, 91, 89,
	3, 2, 2, 2, 92, 93, 7, 8, 2, 2, 93, 17, 3, 2, 2, 2, 94, 100, 7, 16, 2,
	2, 95, 100, 7, 17, 2, 2, 96, 100, 7, 19, 2, 2, 97, 100, 5, 20, 11, 2, 98,
	100, 5, 22, 12, 2, 99, 94, 3, 2, 2, 2, 99, 95, 3, 2, 2, 2, 99, 96, 3, 2,
	2, 2, 99, 97, 3, 2, 2, 2, 99, 98, 3, 2, 2, 2, 100, 19, 3, 2, 2, 2, 101,
	102, 7, 18, 2, 2, 102, 103, 7, 3, 2, 2, 103, 104, 5, 24, 13, 2, 104, 105,
	7, 4, 2, 2, 105, 21, 3, 2, 2, 2, 106, 107, 7, 20, 2, 2, 107, 108, 7, 3,
	2, 2, 108, 109, 5, 24, 13, 2, 109, 110, 7, 10, 2, 2, 110, 111, 5, 24, 13,
	2, 111, 112, 7, 4, 2, 2, 112, 23, 3, 2, 2, 2, 113, 116, 7, 36, 2, 2, 114,
	116, 5, 18, 10, 2, 115, 113, 3, 2, 2, 2, 115, 114, 3, 2, 2, 2, 116, 25,
	3, 2, 2, 2, 117, 118, 7, 24, 2, 2, 118, 119, 7, 9, 2, 2, 119, 120, 7, 36,
	2, 2, 120, 27, 3, 2, 2, 2, 121, 134, 7, 25, 2, 2, 122, 134, 7, 26, 2, 2,
	123, 134, 7, 27, 2, 2, 124, 134, 7, 28, 2, 2, 125, 134, 7, 29, 2, 2, 126,
	134, 7, 30, 2, 2, 127, 134, 7, 31, 2, 2, 128, 134, 7, 32, 2, 2, 129, 134,
	7, 33, 2, 2, 130, 134, 7, 34, 2, 2, 131, 134, 7, 35, 2, 2, 132, 134, 5,
	26, 14, 2, 133, 121, 3, 2, 2, 2, 133, 122, 3, 2, 2, 2, 133, 123, 3, 2,
	2, 2, 133, 124, 3, 2, 2, 2, 133, 125, 3, 2, 2, 2, 133, 126, 3, 2, 2, 2,
	133, 127, 3, 2, 2, 2, 133, 128, 3, 2, 2, 2, 133, 129, 3, 2, 2, 2, 133,
	130, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 133, 132, 3, 2, 2, 2, 134, 29, 3,
	2, 2, 2, 135, 136, 5, 28, 15, 2, 136, 31, 3, 2, 2, 2, 137, 142, 7, 21,
	2, 2, 138, 142, 7, 23, 2, 2, 139, 142, 7, 22, 2, 2, 140, 142, 5, 30, 16,
	2, 141, 137, 3, 2, 2, 2, 141, 138, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2, 141,
	140, 3, 2, 2, 2, 142, 33, 3, 2, 2, 2, 143, 144, 5, 24, 13, 2, 144, 153,
	7, 36, 2, 2, 145, 147, 7, 12, 2, 2, 146, 148, 5, 32, 17, 2, 147, 146, 3,
	2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 149, 150, 3, 2, 2,
	2, 150, 151, 3, 2, 2, 2, 151, 152, 7, 12, 2, 2, 152, 154, 3, 2, 2, 2, 153,
	145, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 35, 3, 2, 2, 2, 155, 156, 5,
	34, 18, 2, 156, 37, 3, 2, 2, 2, 157, 161, 5, 6, 4, 2, 158, 161, 5, 8, 5,
	2, 159, 161, 5, 4, 3, 2, 160, 157, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160,
	159, 3, 2, 2, 2, 161, 39, 3, 2, 2, 2, 162, 164, 5, 38, 20, 2, 163, 162,
	3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 165, 166, 3, 2,
	2, 2, 166, 41, 3, 2, 2, 2, 15, 43, 66, 76, 83, 89, 99, 115, 133, 141, 149,
	153, 160, 165,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'<'", "'>'", "'('", "')'", "'{'", "'}'", "'.'", "','", "'_'", "'`'",
	"'base'", "'comp'", "'enum'", "", "", "", "", "'map'", "", "", "", "",
	"'format.password'", "'format.date'", "'format.datetime'", "'format.byte'",
	"'format.binary'", "'format.email'", "'format.uuid'", "'format.uri'", "'format.hostname'",
	"'format.ipv4'", "'format.ipv6'",
}
var symbolicNames = []string{
	"", "LessThan", "MoreThan", "OpenParen", "CloseParen", "OpenBrace", "CloseBrace",
	"Dot", "Comma", "Underscore", "Tick", "BaseInit", "ComponentInit", "EnumInit",
	"StringType", "NumberType", "ArrayType", "BooleanType", "MapType", "Required",
	"Searchable", "Indexed", "Format", "PasswordFormat", "DateFormat", "DateTimeFormat",
	"ByteFormat", "BinaryFormat", "EmailFormat", "UuidFormat", "UriFormat",
	"HostnameFormat", "Ipv4Format", "Ipv6Format", "Identifier", "WS", "COMMENT",
	"LINE_COMMENT",
}

var ruleNames = []string{
	"program", "baseDeclaration", "objDeclaration", "enumDeclaration", "enumBlock",
	"identifierList", "extendsStatement", "objTail", "nativeType", "arrayType",
	"mapType", "typeName", "formatIdentifier", "formatType", "formatAnnotation",
	"annotationType", "parameterDefinition", "objElement", "modelElement",
	"modelElements",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SwaggableParser struct {
	*antlr.BaseParser
}

func NewSwaggableParser(input antlr.TokenStream) *SwaggableParser {
	this := new(SwaggableParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "SwaggableParser.g4"

	return this
}

// SwaggableParser tokens.
const (
	SwaggableParserEOF            = antlr.TokenEOF
	SwaggableParserLessThan       = 1
	SwaggableParserMoreThan       = 2
	SwaggableParserOpenParen      = 3
	SwaggableParserCloseParen     = 4
	SwaggableParserOpenBrace      = 5
	SwaggableParserCloseBrace     = 6
	SwaggableParserDot            = 7
	SwaggableParserComma          = 8
	SwaggableParserUnderscore     = 9
	SwaggableParserTick           = 10
	SwaggableParserBaseInit       = 11
	SwaggableParserComponentInit  = 12
	SwaggableParserEnumInit       = 13
	SwaggableParserStringType     = 14
	SwaggableParserNumberType     = 15
	SwaggableParserArrayType      = 16
	SwaggableParserBooleanType    = 17
	SwaggableParserMapType        = 18
	SwaggableParserRequired       = 19
	SwaggableParserSearchable     = 20
	SwaggableParserIndexed        = 21
	SwaggableParserFormat         = 22
	SwaggableParserPasswordFormat = 23
	SwaggableParserDateFormat     = 24
	SwaggableParserDateTimeFormat = 25
	SwaggableParserByteFormat     = 26
	SwaggableParserBinaryFormat   = 27
	SwaggableParserEmailFormat    = 28
	SwaggableParserUuidFormat     = 29
	SwaggableParserUriFormat      = 30
	SwaggableParserHostnameFormat = 31
	SwaggableParserIpv4Format     = 32
	SwaggableParserIpv6Format     = 33
	SwaggableParserIdentifier     = 34
	SwaggableParserWS             = 35
	SwaggableParserCOMMENT        = 36
	SwaggableParserLINE_COMMENT   = 37
)

// SwaggableParser rules.
const (
	SwaggableParserRULE_program             = 0
	SwaggableParserRULE_baseDeclaration     = 1
	SwaggableParserRULE_objDeclaration      = 2
	SwaggableParserRULE_enumDeclaration     = 3
	SwaggableParserRULE_enumBlock           = 4
	SwaggableParserRULE_identifierList      = 5
	SwaggableParserRULE_extendsStatement    = 6
	SwaggableParserRULE_objTail             = 7
	SwaggableParserRULE_nativeType          = 8
	SwaggableParserRULE_arrayType           = 9
	SwaggableParserRULE_mapType             = 10
	SwaggableParserRULE_typeName            = 11
	SwaggableParserRULE_formatIdentifier    = 12
	SwaggableParserRULE_formatType          = 13
	SwaggableParserRULE_formatAnnotation    = 14
	SwaggableParserRULE_annotationType      = 15
	SwaggableParserRULE_parameterDefinition = 16
	SwaggableParserRULE_objElement          = 17
	SwaggableParserRULE_modelElement        = 18
	SwaggableParserRULE_modelElements       = 19
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(SwaggableParserEOF, 0)
}

func (s *ProgramContext) ModelElements() IModelElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModelElementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModelElementsContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *SwaggableParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SwaggableParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SwaggableParserBaseInit)|(1<<SwaggableParserComponentInit)|(1<<SwaggableParserEnumInit))) != 0 {
		{
			p.SetState(40)
			p.ModelElements()
		}

	}
	{
		p.SetState(43)
		p.Match(SwaggableParserEOF)
	}

	return localctx
}

// IBaseDeclarationContext is an interface to support dynamic dispatch.
type IBaseDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBaseDeclarationContext differentiates from other interfaces.
	IsBaseDeclarationContext()
}

type BaseDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseDeclarationContext() *BaseDeclarationContext {
	var p = new(BaseDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_baseDeclaration
	return p
}

func (*BaseDeclarationContext) IsBaseDeclarationContext() {}

func NewBaseDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseDeclarationContext {
	var p = new(BaseDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_baseDeclaration

	return p
}

func (s *BaseDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseDeclarationContext) BaseInit() antlr.TerminalNode {
	return s.GetToken(SwaggableParserBaseInit, 0)
}

func (s *BaseDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SwaggableParserIdentifier, 0)
}

func (s *BaseDeclarationContext) ObjTail() IObjTailContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjTailContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjTailContext)
}

func (s *BaseDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterBaseDeclaration(s)
	}
}

func (s *BaseDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitBaseDeclaration(s)
	}
}

func (p *SwaggableParser) BaseDeclaration() (localctx IBaseDeclarationContext) {
	localctx = NewBaseDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SwaggableParserRULE_baseDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(SwaggableParserBaseInit)
	}
	{
		p.SetState(46)
		p.Match(SwaggableParserIdentifier)
	}
	{
		p.SetState(47)
		p.ObjTail()
	}

	return localctx
}

// IObjDeclarationContext is an interface to support dynamic dispatch.
type IObjDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjDeclarationContext differentiates from other interfaces.
	IsObjDeclarationContext()
}

type ObjDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjDeclarationContext() *ObjDeclarationContext {
	var p = new(ObjDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_objDeclaration
	return p
}

func (*ObjDeclarationContext) IsObjDeclarationContext() {}

func NewObjDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjDeclarationContext {
	var p = new(ObjDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_objDeclaration

	return p
}

func (s *ObjDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjDeclarationContext) ComponentInit() antlr.TerminalNode {
	return s.GetToken(SwaggableParserComponentInit, 0)
}

func (s *ObjDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SwaggableParserIdentifier, 0)
}

func (s *ObjDeclarationContext) ObjTail() IObjTailContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjTailContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjTailContext)
}

func (s *ObjDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterObjDeclaration(s)
	}
}

func (s *ObjDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitObjDeclaration(s)
	}
}

func (p *SwaggableParser) ObjDeclaration() (localctx IObjDeclarationContext) {
	localctx = NewObjDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SwaggableParserRULE_objDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(SwaggableParserComponentInit)
	}
	{
		p.SetState(50)
		p.Match(SwaggableParserIdentifier)
	}
	{
		p.SetState(51)
		p.ObjTail()
	}

	return localctx
}

// IEnumDeclarationContext is an interface to support dynamic dispatch.
type IEnumDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumDeclarationContext differentiates from other interfaces.
	IsEnumDeclarationContext()
}

type EnumDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDeclarationContext() *EnumDeclarationContext {
	var p = new(EnumDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_enumDeclaration
	return p
}

func (*EnumDeclarationContext) IsEnumDeclarationContext() {}

func NewEnumDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDeclarationContext {
	var p = new(EnumDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_enumDeclaration

	return p
}

func (s *EnumDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDeclarationContext) EnumInit() antlr.TerminalNode {
	return s.GetToken(SwaggableParserEnumInit, 0)
}

func (s *EnumDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SwaggableParserIdentifier, 0)
}

func (s *EnumDeclarationContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *EnumDeclarationContext) EnumBlock() IEnumBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumBlockContext)
}

func (s *EnumDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterEnumDeclaration(s)
	}
}

func (s *EnumDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitEnumDeclaration(s)
	}
}

func (p *SwaggableParser) EnumDeclaration() (localctx IEnumDeclarationContext) {
	localctx = NewEnumDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwaggableParserRULE_enumDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(SwaggableParserEnumInit)
	}
	{
		p.SetState(54)
		p.Match(SwaggableParserIdentifier)
	}
	{
		p.SetState(55)
		p.Match(SwaggableParserLessThan)
	}
	{
		p.SetState(56)
		p.TypeName()
	}
	{
		p.SetState(57)
		p.Match(SwaggableParserMoreThan)
	}
	{
		p.SetState(58)
		p.EnumBlock()
	}

	return localctx
}

// IEnumBlockContext is an interface to support dynamic dispatch.
type IEnumBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumBlockContext differentiates from other interfaces.
	IsEnumBlockContext()
}

type EnumBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumBlockContext() *EnumBlockContext {
	var p = new(EnumBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_enumBlock
	return p
}

func (*EnumBlockContext) IsEnumBlockContext() {}

func NewEnumBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBlockContext {
	var p = new(EnumBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_enumBlock

	return p
}

func (s *EnumBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBlockContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(SwaggableParserIdentifier)
}

func (s *EnumBlockContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(SwaggableParserIdentifier, i)
}

func (s *EnumBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterEnumBlock(s)
	}
}

func (s *EnumBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitEnumBlock(s)
	}
}

func (p *SwaggableParser) EnumBlock() (localctx IEnumBlockContext) {
	localctx = NewEnumBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SwaggableParserRULE_enumBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(SwaggableParserOpenBrace)
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SwaggableParserIdentifier {
		{
			p.SetState(61)
			p.Match(SwaggableParserIdentifier)
		}

		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(67)
		p.Match(SwaggableParserCloseBrace)
	}

	return localctx
}

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_identifierList
	return p
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(SwaggableParserIdentifier)
}

func (s *IdentifierListContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(SwaggableParserIdentifier, i)
}

func (s *IdentifierListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(SwaggableParserComma)
}

func (s *IdentifierListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(SwaggableParserComma, i)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterIdentifierList(s)
	}
}

func (s *IdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitIdentifierList(s)
	}
}

func (p *SwaggableParser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SwaggableParserRULE_identifierList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(SwaggableParserIdentifier)
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SwaggableParserComma {
		{
			p.SetState(70)
			p.Match(SwaggableParserComma)
		}
		{
			p.SetState(71)
			p.Match(SwaggableParserIdentifier)
		}

		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExtendsStatementContext is an interface to support dynamic dispatch.
type IExtendsStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtendsStatementContext differentiates from other interfaces.
	IsExtendsStatementContext()
}

type ExtendsStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendsStatementContext() *ExtendsStatementContext {
	var p = new(ExtendsStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_extendsStatement
	return p
}

func (*ExtendsStatementContext) IsExtendsStatementContext() {}

func NewExtendsStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendsStatementContext {
	var p = new(ExtendsStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_extendsStatement

	return p
}

func (s *ExtendsStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendsStatementContext) LessThan() antlr.TerminalNode {
	return s.GetToken(SwaggableParserLessThan, 0)
}

func (s *ExtendsStatementContext) IdentifierList() IIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *ExtendsStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendsStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendsStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterExtendsStatement(s)
	}
}

func (s *ExtendsStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitExtendsStatement(s)
	}
}

func (p *SwaggableParser) ExtendsStatement() (localctx IExtendsStatementContext) {
	localctx = NewExtendsStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SwaggableParserRULE_extendsStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(SwaggableParserLessThan)
	}
	{
		p.SetState(78)
		p.IdentifierList()
	}

	return localctx
}

// IObjTailContext is an interface to support dynamic dispatch.
type IObjTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjTailContext differentiates from other interfaces.
	IsObjTailContext()
}

type ObjTailContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjTailContext() *ObjTailContext {
	var p = new(ObjTailContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_objTail
	return p
}

func (*ObjTailContext) IsObjTailContext() {}

func NewObjTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjTailContext {
	var p = new(ObjTailContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_objTail

	return p
}

func (s *ObjTailContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjTailContext) ExtendsStatement() IExtendsStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtendsStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtendsStatementContext)
}

func (s *ObjTailContext) AllObjElement() []IObjElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObjElementContext)(nil)).Elem())
	var tst = make([]IObjElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObjElementContext)
		}
	}

	return tst
}

func (s *ObjTailContext) ObjElement(i int) IObjElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObjElementContext)
}

func (s *ObjTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjTailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterObjTail(s)
	}
}

func (s *ObjTailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitObjTail(s)
	}
}

func (p *SwaggableParser) ObjTail() (localctx IObjTailContext) {
	localctx = NewObjTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SwaggableParserRULE_objTail)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SwaggableParserLessThan {
		{
			p.SetState(80)
			p.ExtendsStatement()
		}

	}
	{
		p.SetState(83)
		p.Match(SwaggableParserOpenBrace)
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-14)&-(0x1f+1)) == 0 && ((1<<uint((_la-14)))&((1<<(SwaggableParserStringType-14))|(1<<(SwaggableParserNumberType-14))|(1<<(SwaggableParserArrayType-14))|(1<<(SwaggableParserBooleanType-14))|(1<<(SwaggableParserMapType-14))|(1<<(SwaggableParserIdentifier-14)))) != 0 {
		{
			p.SetState(84)
			p.ObjElement()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.Match(SwaggableParserCloseBrace)
	}

	return localctx
}

// INativeTypeContext is an interface to support dynamic dispatch.
type INativeTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNativeTypeContext differentiates from other interfaces.
	IsNativeTypeContext()
}

type NativeTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNativeTypeContext() *NativeTypeContext {
	var p = new(NativeTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_nativeType
	return p
}

func (*NativeTypeContext) IsNativeTypeContext() {}

func NewNativeTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NativeTypeContext {
	var p = new(NativeTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_nativeType

	return p
}

func (s *NativeTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *NativeTypeContext) StringType() antlr.TerminalNode {
	return s.GetToken(SwaggableParserStringType, 0)
}

func (s *NativeTypeContext) NumberType() antlr.TerminalNode {
	return s.GetToken(SwaggableParserNumberType, 0)
}

func (s *NativeTypeContext) BooleanType() antlr.TerminalNode {
	return s.GetToken(SwaggableParserBooleanType, 0)
}

func (s *NativeTypeContext) ArrayType() IArrayTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *NativeTypeContext) MapType() IMapTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapTypeContext)
}

func (s *NativeTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NativeTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NativeTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterNativeType(s)
	}
}

func (s *NativeTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitNativeType(s)
	}
}

func (p *SwaggableParser) NativeType() (localctx INativeTypeContext) {
	localctx = NewNativeTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SwaggableParserRULE_nativeType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(97)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SwaggableParserStringType:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Match(SwaggableParserStringType)
		}

	case SwaggableParserNumberType:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Match(SwaggableParserNumberType)
		}

	case SwaggableParserBooleanType:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(94)
			p.Match(SwaggableParserBooleanType)
		}

	case SwaggableParserArrayType:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(95)
			p.ArrayType()
		}

	case SwaggableParserMapType:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(96)
			p.MapType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArrayTypeContext is an interface to support dynamic dispatch.
type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayTypeContext() *ArrayTypeContext {
	var p = new(ArrayTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_arrayType
	return p
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_arrayType

	return p
}

func (s *ArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeContext) ArrayType() antlr.TerminalNode {
	return s.GetToken(SwaggableParserArrayType, 0)
}

func (s *ArrayTypeContext) LessThan() antlr.TerminalNode {
	return s.GetToken(SwaggableParserLessThan, 0)
}

func (s *ArrayTypeContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *ArrayTypeContext) MoreThan() antlr.TerminalNode {
	return s.GetToken(SwaggableParserMoreThan, 0)
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterArrayType(s)
	}
}

func (s *ArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitArrayType(s)
	}
}

func (p *SwaggableParser) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SwaggableParserRULE_arrayType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(SwaggableParserArrayType)
	}
	{
		p.SetState(100)
		p.Match(SwaggableParserLessThan)
	}
	{
		p.SetState(101)
		p.TypeName()
	}
	{
		p.SetState(102)
		p.Match(SwaggableParserMoreThan)
	}

	return localctx
}

// IMapTypeContext is an interface to support dynamic dispatch.
type IMapTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapTypeContext differentiates from other interfaces.
	IsMapTypeContext()
}

type MapTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapTypeContext() *MapTypeContext {
	var p = new(MapTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_mapType
	return p
}

func (*MapTypeContext) IsMapTypeContext() {}

func NewMapTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTypeContext {
	var p = new(MapTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_mapType

	return p
}

func (s *MapTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MapTypeContext) MapType() antlr.TerminalNode {
	return s.GetToken(SwaggableParserMapType, 0)
}

func (s *MapTypeContext) LessThan() antlr.TerminalNode {
	return s.GetToken(SwaggableParserLessThan, 0)
}

func (s *MapTypeContext) AllTypeName() []ITypeNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeNameContext)(nil)).Elem())
	var tst = make([]ITypeNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeNameContext)
		}
	}

	return tst
}

func (s *MapTypeContext) TypeName(i int) ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *MapTypeContext) Comma() antlr.TerminalNode {
	return s.GetToken(SwaggableParserComma, 0)
}

func (s *MapTypeContext) MoreThan() antlr.TerminalNode {
	return s.GetToken(SwaggableParserMoreThan, 0)
}

func (s *MapTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterMapType(s)
	}
}

func (s *MapTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitMapType(s)
	}
}

func (p *SwaggableParser) MapType() (localctx IMapTypeContext) {
	localctx = NewMapTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SwaggableParserRULE_mapType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(SwaggableParserMapType)
	}
	{
		p.SetState(105)
		p.Match(SwaggableParserLessThan)
	}
	{
		p.SetState(106)
		p.TypeName()
	}
	{
		p.SetState(107)
		p.Match(SwaggableParserComma)
	}
	{
		p.SetState(108)
		p.TypeName()
	}
	{
		p.SetState(109)
		p.Match(SwaggableParserMoreThan)
	}

	return localctx
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_typeName
	return p
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SwaggableParserIdentifier, 0)
}

func (s *TypeNameContext) NativeType() INativeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INativeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INativeTypeContext)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (p *SwaggableParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SwaggableParserRULE_typeName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SwaggableParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.Match(SwaggableParserIdentifier)
		}

	case SwaggableParserStringType, SwaggableParserNumberType, SwaggableParserArrayType, SwaggableParserBooleanType, SwaggableParserMapType:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.NativeType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFormatIdentifierContext is an interface to support dynamic dispatch.
type IFormatIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormatIdentifierContext differentiates from other interfaces.
	IsFormatIdentifierContext()
}

type FormatIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormatIdentifierContext() *FormatIdentifierContext {
	var p = new(FormatIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_formatIdentifier
	return p
}

func (*FormatIdentifierContext) IsFormatIdentifierContext() {}

func NewFormatIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormatIdentifierContext {
	var p = new(FormatIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_formatIdentifier

	return p
}

func (s *FormatIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *FormatIdentifierContext) Format() antlr.TerminalNode {
	return s.GetToken(SwaggableParserFormat, 0)
}

func (s *FormatIdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SwaggableParserIdentifier, 0)
}

func (s *FormatIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormatIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormatIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterFormatIdentifier(s)
	}
}

func (s *FormatIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitFormatIdentifier(s)
	}
}

func (p *SwaggableParser) FormatIdentifier() (localctx IFormatIdentifierContext) {
	localctx = NewFormatIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SwaggableParserRULE_formatIdentifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(SwaggableParserFormat)
	}
	{
		p.SetState(116)
		p.Match(SwaggableParserDot)
	}
	{
		p.SetState(117)
		p.Match(SwaggableParserIdentifier)
	}

	return localctx
}

// IFormatTypeContext is an interface to support dynamic dispatch.
type IFormatTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormatTypeContext differentiates from other interfaces.
	IsFormatTypeContext()
}

type FormatTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormatTypeContext() *FormatTypeContext {
	var p = new(FormatTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_formatType
	return p
}

func (*FormatTypeContext) IsFormatTypeContext() {}

func NewFormatTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormatTypeContext {
	var p = new(FormatTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_formatType

	return p
}

func (s *FormatTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FormatTypeContext) PasswordFormat() antlr.TerminalNode {
	return s.GetToken(SwaggableParserPasswordFormat, 0)
}

func (s *FormatTypeContext) DateFormat() antlr.TerminalNode {
	return s.GetToken(SwaggableParserDateFormat, 0)
}

func (s *FormatTypeContext) DateTimeFormat() antlr.TerminalNode {
	return s.GetToken(SwaggableParserDateTimeFormat, 0)
}

func (s *FormatTypeContext) ByteFormat() antlr.TerminalNode {
	return s.GetToken(SwaggableParserByteFormat, 0)
}

func (s *FormatTypeContext) BinaryFormat() antlr.TerminalNode {
	return s.GetToken(SwaggableParserBinaryFormat, 0)
}

func (s *FormatTypeContext) EmailFormat() antlr.TerminalNode {
	return s.GetToken(SwaggableParserEmailFormat, 0)
}

func (s *FormatTypeContext) UuidFormat() antlr.TerminalNode {
	return s.GetToken(SwaggableParserUuidFormat, 0)
}

func (s *FormatTypeContext) UriFormat() antlr.TerminalNode {
	return s.GetToken(SwaggableParserUriFormat, 0)
}

func (s *FormatTypeContext) HostnameFormat() antlr.TerminalNode {
	return s.GetToken(SwaggableParserHostnameFormat, 0)
}

func (s *FormatTypeContext) Ipv4Format() antlr.TerminalNode {
	return s.GetToken(SwaggableParserIpv4Format, 0)
}

func (s *FormatTypeContext) Ipv6Format() antlr.TerminalNode {
	return s.GetToken(SwaggableParserIpv6Format, 0)
}

func (s *FormatTypeContext) FormatIdentifier() IFormatIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormatIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormatIdentifierContext)
}

func (s *FormatTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormatTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormatTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterFormatType(s)
	}
}

func (s *FormatTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitFormatType(s)
	}
}

func (p *SwaggableParser) FormatType() (localctx IFormatTypeContext) {
	localctx = NewFormatTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SwaggableParserRULE_formatType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SwaggableParserPasswordFormat:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Match(SwaggableParserPasswordFormat)
		}

	case SwaggableParserDateFormat:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Match(SwaggableParserDateFormat)
		}

	case SwaggableParserDateTimeFormat:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(121)
			p.Match(SwaggableParserDateTimeFormat)
		}

	case SwaggableParserByteFormat:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(122)
			p.Match(SwaggableParserByteFormat)
		}

	case SwaggableParserBinaryFormat:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(123)
			p.Match(SwaggableParserBinaryFormat)
		}

	case SwaggableParserEmailFormat:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(124)
			p.Match(SwaggableParserEmailFormat)
		}

	case SwaggableParserUuidFormat:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(125)
			p.Match(SwaggableParserUuidFormat)
		}

	case SwaggableParserUriFormat:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(126)
			p.Match(SwaggableParserUriFormat)
		}

	case SwaggableParserHostnameFormat:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(127)
			p.Match(SwaggableParserHostnameFormat)
		}

	case SwaggableParserIpv4Format:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(128)
			p.Match(SwaggableParserIpv4Format)
		}

	case SwaggableParserIpv6Format:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(129)
			p.Match(SwaggableParserIpv6Format)
		}

	case SwaggableParserFormat:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(130)
			p.FormatIdentifier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFormatAnnotationContext is an interface to support dynamic dispatch.
type IFormatAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormatAnnotationContext differentiates from other interfaces.
	IsFormatAnnotationContext()
}

type FormatAnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormatAnnotationContext() *FormatAnnotationContext {
	var p = new(FormatAnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_formatAnnotation
	return p
}

func (*FormatAnnotationContext) IsFormatAnnotationContext() {}

func NewFormatAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormatAnnotationContext {
	var p = new(FormatAnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_formatAnnotation

	return p
}

func (s *FormatAnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *FormatAnnotationContext) FormatType() IFormatTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormatTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormatTypeContext)
}

func (s *FormatAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormatAnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormatAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterFormatAnnotation(s)
	}
}

func (s *FormatAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitFormatAnnotation(s)
	}
}

func (p *SwaggableParser) FormatAnnotation() (localctx IFormatAnnotationContext) {
	localctx = NewFormatAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SwaggableParserRULE_formatAnnotation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.FormatType()
	}

	return localctx
}

// IAnnotationTypeContext is an interface to support dynamic dispatch.
type IAnnotationTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotationTypeContext differentiates from other interfaces.
	IsAnnotationTypeContext()
}

type AnnotationTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationTypeContext() *AnnotationTypeContext {
	var p = new(AnnotationTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_annotationType
	return p
}

func (*AnnotationTypeContext) IsAnnotationTypeContext() {}

func NewAnnotationTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationTypeContext {
	var p = new(AnnotationTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_annotationType

	return p
}

func (s *AnnotationTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationTypeContext) Required() antlr.TerminalNode {
	return s.GetToken(SwaggableParserRequired, 0)
}

func (s *AnnotationTypeContext) Indexed() antlr.TerminalNode {
	return s.GetToken(SwaggableParserIndexed, 0)
}

func (s *AnnotationTypeContext) Searchable() antlr.TerminalNode {
	return s.GetToken(SwaggableParserSearchable, 0)
}

func (s *AnnotationTypeContext) FormatAnnotation() IFormatAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormatAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormatAnnotationContext)
}

func (s *AnnotationTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterAnnotationType(s)
	}
}

func (s *AnnotationTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitAnnotationType(s)
	}
}

func (p *SwaggableParser) AnnotationType() (localctx IAnnotationTypeContext) {
	localctx = NewAnnotationTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SwaggableParserRULE_annotationType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(139)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SwaggableParserRequired:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.Match(SwaggableParserRequired)
		}

	case SwaggableParserIndexed:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)
			p.Match(SwaggableParserIndexed)
		}

	case SwaggableParserSearchable:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(137)
			p.Match(SwaggableParserSearchable)
		}

	case SwaggableParserFormat, SwaggableParserPasswordFormat, SwaggableParserDateFormat, SwaggableParserDateTimeFormat, SwaggableParserByteFormat, SwaggableParserBinaryFormat, SwaggableParserEmailFormat, SwaggableParserUuidFormat, SwaggableParserUriFormat, SwaggableParserHostnameFormat, SwaggableParserIpv4Format, SwaggableParserIpv6Format:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(138)
			p.FormatAnnotation()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IParameterDefinitionContext is an interface to support dynamic dispatch.
type IParameterDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterDefinitionContext differentiates from other interfaces.
	IsParameterDefinitionContext()
}

type ParameterDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterDefinitionContext() *ParameterDefinitionContext {
	var p = new(ParameterDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_parameterDefinition
	return p
}

func (*ParameterDefinitionContext) IsParameterDefinitionContext() {}

func NewParameterDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterDefinitionContext {
	var p = new(ParameterDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_parameterDefinition

	return p
}

func (s *ParameterDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterDefinitionContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *ParameterDefinitionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(SwaggableParserIdentifier, 0)
}

func (s *ParameterDefinitionContext) AllTick() []antlr.TerminalNode {
	return s.GetTokens(SwaggableParserTick)
}

func (s *ParameterDefinitionContext) Tick(i int) antlr.TerminalNode {
	return s.GetToken(SwaggableParserTick, i)
}

func (s *ParameterDefinitionContext) AllAnnotationType() []IAnnotationTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationTypeContext)(nil)).Elem())
	var tst = make([]IAnnotationTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationTypeContext)
		}
	}

	return tst
}

func (s *ParameterDefinitionContext) AnnotationType(i int) IAnnotationTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationTypeContext)
}

func (s *ParameterDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterParameterDefinition(s)
	}
}

func (s *ParameterDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitParameterDefinition(s)
	}
}

func (p *SwaggableParser) ParameterDefinition() (localctx IParameterDefinitionContext) {
	localctx = NewParameterDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SwaggableParserRULE_parameterDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.TypeName()
	}
	{
		p.SetState(142)
		p.Match(SwaggableParserIdentifier)
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SwaggableParserTick {
		{
			p.SetState(143)
			p.Match(SwaggableParserTick)
		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la-19)&-(0x1f+1)) == 0 && ((1<<uint((_la-19)))&((1<<(SwaggableParserRequired-19))|(1<<(SwaggableParserSearchable-19))|(1<<(SwaggableParserIndexed-19))|(1<<(SwaggableParserFormat-19))|(1<<(SwaggableParserPasswordFormat-19))|(1<<(SwaggableParserDateFormat-19))|(1<<(SwaggableParserDateTimeFormat-19))|(1<<(SwaggableParserByteFormat-19))|(1<<(SwaggableParserBinaryFormat-19))|(1<<(SwaggableParserEmailFormat-19))|(1<<(SwaggableParserUuidFormat-19))|(1<<(SwaggableParserUriFormat-19))|(1<<(SwaggableParserHostnameFormat-19))|(1<<(SwaggableParserIpv4Format-19))|(1<<(SwaggableParserIpv6Format-19)))) != 0) {
			{
				p.SetState(144)
				p.AnnotationType()
			}

			p.SetState(147)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(149)
			p.Match(SwaggableParserTick)
		}

	}

	return localctx
}

// IObjElementContext is an interface to support dynamic dispatch.
type IObjElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjElementContext differentiates from other interfaces.
	IsObjElementContext()
}

type ObjElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjElementContext() *ObjElementContext {
	var p = new(ObjElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_objElement
	return p
}

func (*ObjElementContext) IsObjElementContext() {}

func NewObjElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjElementContext {
	var p = new(ObjElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_objElement

	return p
}

func (s *ObjElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjElementContext) ParameterDefinition() IParameterDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterDefinitionContext)
}

func (s *ObjElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterObjElement(s)
	}
}

func (s *ObjElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitObjElement(s)
	}
}

func (p *SwaggableParser) ObjElement() (localctx IObjElementContext) {
	localctx = NewObjElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SwaggableParserRULE_objElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.ParameterDefinition()
	}

	return localctx
}

// IModelElementContext is an interface to support dynamic dispatch.
type IModelElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModelElementContext differentiates from other interfaces.
	IsModelElementContext()
}

type ModelElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelElementContext() *ModelElementContext {
	var p = new(ModelElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_modelElement
	return p
}

func (*ModelElementContext) IsModelElementContext() {}

func NewModelElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelElementContext {
	var p = new(ModelElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_modelElement

	return p
}

func (s *ModelElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ModelElementContext) ObjDeclaration() IObjDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjDeclarationContext)
}

func (s *ModelElementContext) EnumDeclaration() IEnumDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumDeclarationContext)
}

func (s *ModelElementContext) BaseDeclaration() IBaseDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseDeclarationContext)
}

func (s *ModelElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModelElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterModelElement(s)
	}
}

func (s *ModelElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitModelElement(s)
	}
}

func (p *SwaggableParser) ModelElement() (localctx IModelElementContext) {
	localctx = NewModelElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SwaggableParserRULE_modelElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(158)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SwaggableParserComponentInit:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.ObjDeclaration()
		}

	case SwaggableParserEnumInit:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.EnumDeclaration()
		}

	case SwaggableParserBaseInit:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(157)
			p.BaseDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IModelElementsContext is an interface to support dynamic dispatch.
type IModelElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModelElementsContext differentiates from other interfaces.
	IsModelElementsContext()
}

type ModelElementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelElementsContext() *ModelElementsContext {
	var p = new(ModelElementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SwaggableParserRULE_modelElements
	return p
}

func (*ModelElementsContext) IsModelElementsContext() {}

func NewModelElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelElementsContext {
	var p = new(ModelElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwaggableParserRULE_modelElements

	return p
}

func (s *ModelElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *ModelElementsContext) AllModelElement() []IModelElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IModelElementContext)(nil)).Elem())
	var tst = make([]IModelElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IModelElementContext)
		}
	}

	return tst
}

func (s *ModelElementsContext) ModelElement(i int) IModelElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModelElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IModelElementContext)
}

func (s *ModelElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModelElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.EnterModelElements(s)
	}
}

func (s *ModelElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwaggableParserListener); ok {
		listenerT.ExitModelElements(s)
	}
}

func (p *SwaggableParser) ModelElements() (localctx IModelElementsContext) {
	localctx = NewModelElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SwaggableParserRULE_modelElements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SwaggableParserBaseInit)|(1<<SwaggableParserComponentInit)|(1<<SwaggableParserEnumInit))) != 0) {
		{
			p.SetState(160)
			p.ModelElement()
		}

		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
