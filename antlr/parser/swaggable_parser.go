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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 37, 161,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 5, 2, 42, 10, 2, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	7, 5, 59, 10, 5, 12, 5, 14, 5, 62, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	7, 6, 69, 10, 6, 12, 6, 14, 6, 72, 11, 6, 3, 7, 3, 7, 3, 7, 3, 8, 5, 8,
	78, 10, 8, 3, 8, 3, 8, 7, 8, 82, 10, 8, 12, 8, 14, 8, 85, 11, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 94, 10, 9, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 5, 12, 110, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 128,
	10, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 136, 10, 16, 3,
	17, 3, 17, 3, 17, 3, 17, 6, 17, 142, 10, 17, 13, 17, 14, 17, 143, 3, 17,
	3, 17, 5, 17, 148, 10, 17, 3, 18, 3, 18, 3, 19, 3, 19, 5, 19, 154, 10,
	19, 3, 20, 6, 20, 157, 10, 20, 13, 20, 14, 20, 158, 3, 20, 2, 2, 21, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 2,
	2, 2, 169, 2, 41, 3, 2, 2, 2, 4, 45, 3, 2, 2, 2, 6, 49, 3, 2, 2, 2, 8,
	56, 3, 2, 2, 2, 10, 65, 3, 2, 2, 2, 12, 73, 3, 2, 2, 2, 14, 77, 3, 2, 2,
	2, 16, 93, 3, 2, 2, 2, 18, 95, 3, 2, 2, 2, 20, 100, 3, 2, 2, 2, 22, 109,
	3, 2, 2, 2, 24, 111, 3, 2, 2, 2, 26, 127, 3, 2, 2, 2, 28, 129, 3, 2, 2,
	2, 30, 135, 3, 2, 2, 2, 32, 137, 3, 2, 2, 2, 34, 149, 3, 2, 2, 2, 36, 153,
	3, 2, 2, 2, 38, 156, 3, 2, 2, 2, 40, 42, 5, 38, 20, 2, 41, 40, 3, 2, 2,
	2, 41, 42, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 44, 7, 2, 2, 3, 44, 3, 3,
	2, 2, 2, 45, 46, 7, 12, 2, 2, 46, 47, 7, 34, 2, 2, 47, 48, 5, 14, 8, 2,
	48, 5, 3, 2, 2, 2, 49, 50, 7, 13, 2, 2, 50, 51, 7, 34, 2, 2, 51, 52, 7,
	3, 2, 2, 52, 53, 5, 22, 12, 2, 53, 54, 7, 4, 2, 2, 54, 55, 5, 8, 5, 2,
	55, 7, 3, 2, 2, 2, 56, 60, 7, 7, 2, 2, 57, 59, 7, 34, 2, 2, 58, 57, 3,
	2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61,
	63, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 63, 64, 7, 8, 2, 2, 64, 9, 3, 2, 2,
	2, 65, 70, 7, 34, 2, 2, 66, 67, 7, 10, 2, 2, 67, 69, 7, 34, 2, 2, 68, 66,
	3, 2, 2, 2, 69, 72, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2,
	71, 11, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 73, 74, 7, 3, 2, 2, 74, 75, 5,
	10, 6, 2, 75, 13, 3, 2, 2, 2, 76, 78, 5, 12, 7, 2, 77, 76, 3, 2, 2, 2,
	77, 78, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 83, 7, 7, 2, 2, 80, 82, 5,
	34, 18, 2, 81, 80, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2,
	83, 84, 3, 2, 2, 2, 84, 86, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 87, 7,
	8, 2, 2, 87, 15, 3, 2, 2, 2, 88, 94, 7, 14, 2, 2, 89, 94, 7, 15, 2, 2,
	90, 94, 7, 17, 2, 2, 91, 94, 5, 18, 10, 2, 92, 94, 5, 20, 11, 2, 93, 88,
	3, 2, 2, 2, 93, 89, 3, 2, 2, 2, 93, 90, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2,
	93, 92, 3, 2, 2, 2, 94, 17, 3, 2, 2, 2, 95, 96, 7, 16, 2, 2, 96, 97, 7,
	3, 2, 2, 97, 98, 5, 22, 12, 2, 98, 99, 7, 4, 2, 2, 99, 19, 3, 2, 2, 2,
	100, 101, 7, 18, 2, 2, 101, 102, 7, 3, 2, 2, 102, 103, 5, 22, 12, 2, 103,
	104, 7, 10, 2, 2, 104, 105, 5, 22, 12, 2, 105, 106, 7, 4, 2, 2, 106, 21,
	3, 2, 2, 2, 107, 110, 7, 34, 2, 2, 108, 110, 5, 16, 9, 2, 109, 107, 3,
	2, 2, 2, 109, 108, 3, 2, 2, 2, 110, 23, 3, 2, 2, 2, 111, 112, 7, 22, 2,
	2, 112, 113, 7, 9, 2, 2, 113, 114, 7, 34, 2, 2, 114, 25, 3, 2, 2, 2, 115,
	128, 7, 23, 2, 2, 116, 128, 7, 24, 2, 2, 117, 128, 7, 25, 2, 2, 118, 128,
	7, 26, 2, 2, 119, 128, 7, 27, 2, 2, 120, 128, 7, 28, 2, 2, 121, 128, 7,
	29, 2, 2, 122, 128, 7, 30, 2, 2, 123, 128, 7, 31, 2, 2, 124, 128, 7, 32,
	2, 2, 125, 128, 7, 33, 2, 2, 126, 128, 5, 24, 13, 2, 127, 115, 3, 2, 2,
	2, 127, 116, 3, 2, 2, 2, 127, 117, 3, 2, 2, 2, 127, 118, 3, 2, 2, 2, 127,
	119, 3, 2, 2, 2, 127, 120, 3, 2, 2, 2, 127, 121, 3, 2, 2, 2, 127, 122,
	3, 2, 2, 2, 127, 123, 3, 2, 2, 2, 127, 124, 3, 2, 2, 2, 127, 125, 3, 2,
	2, 2, 127, 126, 3, 2, 2, 2, 128, 27, 3, 2, 2, 2, 129, 130, 5, 26, 14, 2,
	130, 29, 3, 2, 2, 2, 131, 136, 7, 19, 2, 2, 132, 136, 7, 21, 2, 2, 133,
	136, 7, 20, 2, 2, 134, 136, 5, 28, 15, 2, 135, 131, 3, 2, 2, 2, 135, 132,
	3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 135, 134, 3, 2, 2, 2, 136, 31, 3, 2,
	2, 2, 137, 138, 5, 22, 12, 2, 138, 147, 7, 34, 2, 2, 139, 141, 7, 5, 2,
	2, 140, 142, 5, 30, 16, 2, 141, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2,
	143, 141, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145,
	146, 7, 6, 2, 2, 146, 148, 3, 2, 2, 2, 147, 139, 3, 2, 2, 2, 147, 148,
	3, 2, 2, 2, 148, 33, 3, 2, 2, 2, 149, 150, 5, 32, 17, 2, 150, 35, 3, 2,
	2, 2, 151, 154, 5, 4, 3, 2, 152, 154, 5, 6, 4, 2, 153, 151, 3, 2, 2, 2,
	153, 152, 3, 2, 2, 2, 154, 37, 3, 2, 2, 2, 155, 157, 5, 36, 19, 2, 156,
	155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 158, 159,
	3, 2, 2, 2, 159, 39, 3, 2, 2, 2, 15, 41, 60, 70, 77, 83, 93, 109, 127,
	135, 143, 147, 153, 158,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'<'", "'>'", "'('", "')'", "'{'", "'}'", "'.'", "','", "'_'", "'comp'",
	"'enum'", "", "", "", "", "'map'", "", "", "", "", "'format.password'",
	"'format.date'", "'format.datetime'", "'format.byte'", "'format.binary'",
	"'format.email'", "'format.uuid'", "'format.uri'", "'format.hostname'",
	"'format.ipv4'", "'format.ipv6'",
}
var symbolicNames = []string{
	"", "LessThan", "MoreThan", "OpenParen", "CloseParen", "OpenBrace", "CloseBrace",
	"Dot", "Comma", "Underscore", "ComponentInit", "EnumInit", "StringType",
	"NumberType", "ArrayType", "BooleanType", "MapType", "Required", "Searchable",
	"Indexed", "Format", "PasswordFormat", "DateFormat", "DateTimeFormat",
	"ByteFormat", "BinaryFormat", "EmailFormat", "UuidFormat", "UriFormat",
	"HostnameFormat", "Ipv4Format", "Ipv6Format", "Identifier", "WS", "COMMENT",
	"LINE_COMMENT",
}

var ruleNames = []string{
	"program", "objDeclaration", "enumDeclaration", "enumBlock", "identifierList",
	"extendsStatement", "objTail", "nativeType", "arrayType", "mapType", "typeName",
	"formatIdentifier", "formatType", "formatAnnotation", "annotationType",
	"parameterDefinition", "objElement", "modelElement", "modelElements",
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
	SwaggableParserComponentInit  = 10
	SwaggableParserEnumInit       = 11
	SwaggableParserStringType     = 12
	SwaggableParserNumberType     = 13
	SwaggableParserArrayType      = 14
	SwaggableParserBooleanType    = 15
	SwaggableParserMapType        = 16
	SwaggableParserRequired       = 17
	SwaggableParserSearchable     = 18
	SwaggableParserIndexed        = 19
	SwaggableParserFormat         = 20
	SwaggableParserPasswordFormat = 21
	SwaggableParserDateFormat     = 22
	SwaggableParserDateTimeFormat = 23
	SwaggableParserByteFormat     = 24
	SwaggableParserBinaryFormat   = 25
	SwaggableParserEmailFormat    = 26
	SwaggableParserUuidFormat     = 27
	SwaggableParserUriFormat      = 28
	SwaggableParserHostnameFormat = 29
	SwaggableParserIpv4Format     = 30
	SwaggableParserIpv6Format     = 31
	SwaggableParserIdentifier     = 32
	SwaggableParserWS             = 33
	SwaggableParserCOMMENT        = 34
	SwaggableParserLINE_COMMENT   = 35
)

// SwaggableParser rules.
const (
	SwaggableParserRULE_program             = 0
	SwaggableParserRULE_objDeclaration      = 1
	SwaggableParserRULE_enumDeclaration     = 2
	SwaggableParserRULE_enumBlock           = 3
	SwaggableParserRULE_identifierList      = 4
	SwaggableParserRULE_extendsStatement    = 5
	SwaggableParserRULE_objTail             = 6
	SwaggableParserRULE_nativeType          = 7
	SwaggableParserRULE_arrayType           = 8
	SwaggableParserRULE_mapType             = 9
	SwaggableParserRULE_typeName            = 10
	SwaggableParserRULE_formatIdentifier    = 11
	SwaggableParserRULE_formatType          = 12
	SwaggableParserRULE_formatAnnotation    = 13
	SwaggableParserRULE_annotationType      = 14
	SwaggableParserRULE_parameterDefinition = 15
	SwaggableParserRULE_objElement          = 16
	SwaggableParserRULE_modelElement        = 17
	SwaggableParserRULE_modelElements       = 18
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
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SwaggableParserComponentInit || _la == SwaggableParserEnumInit {
		{
			p.SetState(38)
			p.ModelElements()
		}

	}
	{
		p.SetState(41)
		p.Match(SwaggableParserEOF)
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
	p.EnterRule(localctx, 2, SwaggableParserRULE_objDeclaration)

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
		p.SetState(43)
		p.Match(SwaggableParserComponentInit)
	}
	{
		p.SetState(44)
		p.Match(SwaggableParserIdentifier)
	}
	{
		p.SetState(45)
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
	p.EnterRule(localctx, 4, SwaggableParserRULE_enumDeclaration)

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
		p.SetState(47)
		p.Match(SwaggableParserEnumInit)
	}
	{
		p.SetState(48)
		p.Match(SwaggableParserIdentifier)
	}
	{
		p.SetState(49)
		p.Match(SwaggableParserLessThan)
	}
	{
		p.SetState(50)
		p.TypeName()
	}
	{
		p.SetState(51)
		p.Match(SwaggableParserMoreThan)
	}
	{
		p.SetState(52)
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
	p.EnterRule(localctx, 6, SwaggableParserRULE_enumBlock)
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
		p.SetState(54)
		p.Match(SwaggableParserOpenBrace)
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SwaggableParserIdentifier {
		{
			p.SetState(55)
			p.Match(SwaggableParserIdentifier)
		}

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(61)
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
	p.EnterRule(localctx, 8, SwaggableParserRULE_identifierList)
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
		p.SetState(63)
		p.Match(SwaggableParserIdentifier)
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SwaggableParserComma {
		{
			p.SetState(64)
			p.Match(SwaggableParserComma)
		}
		{
			p.SetState(65)
			p.Match(SwaggableParserIdentifier)
		}

		p.SetState(70)
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
	p.EnterRule(localctx, 10, SwaggableParserRULE_extendsStatement)

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
		p.SetState(71)
		p.Match(SwaggableParserLessThan)
	}
	{
		p.SetState(72)
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
	p.EnterRule(localctx, 12, SwaggableParserRULE_objTail)
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
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SwaggableParserLessThan {
		{
			p.SetState(74)
			p.ExtendsStatement()
		}

	}
	{
		p.SetState(77)
		p.Match(SwaggableParserOpenBrace)
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-12)&-(0x1f+1)) == 0 && ((1<<uint((_la-12)))&((1<<(SwaggableParserStringType-12))|(1<<(SwaggableParserNumberType-12))|(1<<(SwaggableParserArrayType-12))|(1<<(SwaggableParserBooleanType-12))|(1<<(SwaggableParserMapType-12))|(1<<(SwaggableParserIdentifier-12)))) != 0 {
		{
			p.SetState(78)
			p.ObjElement()
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(84)
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
	p.EnterRule(localctx, 14, SwaggableParserRULE_nativeType)

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

	p.SetState(91)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SwaggableParserStringType:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(SwaggableParserStringType)
		}

	case SwaggableParserNumberType:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
			p.Match(SwaggableParserNumberType)
		}

	case SwaggableParserBooleanType:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)
			p.Match(SwaggableParserBooleanType)
		}

	case SwaggableParserArrayType:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(89)
			p.ArrayType()
		}

	case SwaggableParserMapType:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(90)
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

func (s *ArrayTypeContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
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
	p.EnterRule(localctx, 16, SwaggableParserRULE_arrayType)

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
		p.SetState(93)
		p.Match(SwaggableParserArrayType)
	}
	{
		p.SetState(94)
		p.Match(SwaggableParserLessThan)
	}
	{
		p.SetState(95)
		p.TypeName()
	}
	{
		p.SetState(96)
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
	p.EnterRule(localctx, 18, SwaggableParserRULE_mapType)

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
		p.SetState(98)
		p.Match(SwaggableParserMapType)
	}
	{
		p.SetState(99)
		p.Match(SwaggableParserLessThan)
	}
	{
		p.SetState(100)
		p.TypeName()
	}
	{
		p.SetState(101)
		p.Match(SwaggableParserComma)
	}
	{
		p.SetState(102)
		p.TypeName()
	}
	{
		p.SetState(103)
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
	p.EnterRule(localctx, 20, SwaggableParserRULE_typeName)

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

	p.SetState(107)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SwaggableParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.Match(SwaggableParserIdentifier)
		}

	case SwaggableParserStringType, SwaggableParserNumberType, SwaggableParserArrayType, SwaggableParserBooleanType, SwaggableParserMapType:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
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
	p.EnterRule(localctx, 22, SwaggableParserRULE_formatIdentifier)

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
		p.SetState(109)
		p.Match(SwaggableParserFormat)
	}
	{
		p.SetState(110)
		p.Match(SwaggableParserDot)
	}
	{
		p.SetState(111)
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
	p.EnterRule(localctx, 24, SwaggableParserRULE_formatType)

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

	p.SetState(125)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SwaggableParserPasswordFormat:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)
			p.Match(SwaggableParserPasswordFormat)
		}

	case SwaggableParserDateFormat:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
			p.Match(SwaggableParserDateFormat)
		}

	case SwaggableParserDateTimeFormat:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(115)
			p.Match(SwaggableParserDateTimeFormat)
		}

	case SwaggableParserByteFormat:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(116)
			p.Match(SwaggableParserByteFormat)
		}

	case SwaggableParserBinaryFormat:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(117)
			p.Match(SwaggableParserBinaryFormat)
		}

	case SwaggableParserEmailFormat:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(118)
			p.Match(SwaggableParserEmailFormat)
		}

	case SwaggableParserUuidFormat:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(119)
			p.Match(SwaggableParserUuidFormat)
		}

	case SwaggableParserUriFormat:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(120)
			p.Match(SwaggableParserUriFormat)
		}

	case SwaggableParserHostnameFormat:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(121)
			p.Match(SwaggableParserHostnameFormat)
		}

	case SwaggableParserIpv4Format:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(122)
			p.Match(SwaggableParserIpv4Format)
		}

	case SwaggableParserIpv6Format:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(123)
			p.Match(SwaggableParserIpv6Format)
		}

	case SwaggableParserFormat:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(124)
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
	p.EnterRule(localctx, 26, SwaggableParserRULE_formatAnnotation)

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
		p.SetState(127)
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
	p.EnterRule(localctx, 28, SwaggableParserRULE_annotationType)

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

	p.SetState(133)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SwaggableParserRequired:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)
			p.Match(SwaggableParserRequired)
		}

	case SwaggableParserIndexed:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.Match(SwaggableParserIndexed)
		}

	case SwaggableParserSearchable:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(131)
			p.Match(SwaggableParserSearchable)
		}

	case SwaggableParserFormat, SwaggableParserPasswordFormat, SwaggableParserDateFormat, SwaggableParserDateTimeFormat, SwaggableParserByteFormat, SwaggableParserBinaryFormat, SwaggableParserEmailFormat, SwaggableParserUuidFormat, SwaggableParserUriFormat, SwaggableParserHostnameFormat, SwaggableParserIpv4Format, SwaggableParserIpv6Format:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(132)
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
	p.EnterRule(localctx, 30, SwaggableParserRULE_parameterDefinition)
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
		p.SetState(135)
		p.TypeName()
	}
	{
		p.SetState(136)
		p.Match(SwaggableParserIdentifier)
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SwaggableParserOpenParen {
		{
			p.SetState(137)
			p.Match(SwaggableParserOpenParen)
		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SwaggableParserRequired)|(1<<SwaggableParserSearchable)|(1<<SwaggableParserIndexed)|(1<<SwaggableParserFormat)|(1<<SwaggableParserPasswordFormat)|(1<<SwaggableParserDateFormat)|(1<<SwaggableParserDateTimeFormat)|(1<<SwaggableParserByteFormat)|(1<<SwaggableParserBinaryFormat)|(1<<SwaggableParserEmailFormat)|(1<<SwaggableParserUuidFormat)|(1<<SwaggableParserUriFormat)|(1<<SwaggableParserHostnameFormat)|(1<<SwaggableParserIpv4Format)|(1<<SwaggableParserIpv6Format))) != 0) {
			{
				p.SetState(138)
				p.AnnotationType()
			}

			p.SetState(141)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(143)
			p.Match(SwaggableParserCloseParen)
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
	p.EnterRule(localctx, 32, SwaggableParserRULE_objElement)

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
		p.SetState(147)
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
	p.EnterRule(localctx, 34, SwaggableParserRULE_modelElement)

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

	p.SetState(151)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SwaggableParserComponentInit:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.ObjDeclaration()
		}

	case SwaggableParserEnumInit:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.EnumDeclaration()
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
	p.EnterRule(localctx, 36, SwaggableParserRULE_modelElements)
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
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SwaggableParserComponentInit || _la == SwaggableParserEnumInit {
		{
			p.SetState(153)
			p.ModelElement()
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
