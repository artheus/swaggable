// Code generated from SwaggableParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // SwaggableParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSwaggableParserListener is a complete listener for a parse tree produced by SwaggableParser.
type BaseSwaggableParserListener struct{}

var _ SwaggableParserListener = &BaseSwaggableParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSwaggableParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSwaggableParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSwaggableParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSwaggableParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseSwaggableParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseSwaggableParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterObjDeclaration is called when production objDeclaration is entered.
func (s *BaseSwaggableParserListener) EnterObjDeclaration(ctx *ObjDeclarationContext) {}

// ExitObjDeclaration is called when production objDeclaration is exited.
func (s *BaseSwaggableParserListener) ExitObjDeclaration(ctx *ObjDeclarationContext) {}

// EnterEnumDeclaration is called when production enumDeclaration is entered.
func (s *BaseSwaggableParserListener) EnterEnumDeclaration(ctx *EnumDeclarationContext) {}

// ExitEnumDeclaration is called when production enumDeclaration is exited.
func (s *BaseSwaggableParserListener) ExitEnumDeclaration(ctx *EnumDeclarationContext) {}

// EnterEnumBlock is called when production enumBlock is entered.
func (s *BaseSwaggableParserListener) EnterEnumBlock(ctx *EnumBlockContext) {}

// ExitEnumBlock is called when production enumBlock is exited.
func (s *BaseSwaggableParserListener) ExitEnumBlock(ctx *EnumBlockContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseSwaggableParserListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseSwaggableParserListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterExtendsStatement is called when production extendsStatement is entered.
func (s *BaseSwaggableParserListener) EnterExtendsStatement(ctx *ExtendsStatementContext) {}

// ExitExtendsStatement is called when production extendsStatement is exited.
func (s *BaseSwaggableParserListener) ExitExtendsStatement(ctx *ExtendsStatementContext) {}

// EnterObjTail is called when production objTail is entered.
func (s *BaseSwaggableParserListener) EnterObjTail(ctx *ObjTailContext) {}

// ExitObjTail is called when production objTail is exited.
func (s *BaseSwaggableParserListener) ExitObjTail(ctx *ObjTailContext) {}

// EnterNativeType is called when production nativeType is entered.
func (s *BaseSwaggableParserListener) EnterNativeType(ctx *NativeTypeContext) {}

// ExitNativeType is called when production nativeType is exited.
func (s *BaseSwaggableParserListener) ExitNativeType(ctx *NativeTypeContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseSwaggableParserListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseSwaggableParserListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseSwaggableParserListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseSwaggableParserListener) ExitMapType(ctx *MapTypeContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseSwaggableParserListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseSwaggableParserListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterFormatIdentifier is called when production formatIdentifier is entered.
func (s *BaseSwaggableParserListener) EnterFormatIdentifier(ctx *FormatIdentifierContext) {}

// ExitFormatIdentifier is called when production formatIdentifier is exited.
func (s *BaseSwaggableParserListener) ExitFormatIdentifier(ctx *FormatIdentifierContext) {}

// EnterFormatType is called when production formatType is entered.
func (s *BaseSwaggableParserListener) EnterFormatType(ctx *FormatTypeContext) {}

// ExitFormatType is called when production formatType is exited.
func (s *BaseSwaggableParserListener) ExitFormatType(ctx *FormatTypeContext) {}

// EnterFormatAnnotation is called when production formatAnnotation is entered.
func (s *BaseSwaggableParserListener) EnterFormatAnnotation(ctx *FormatAnnotationContext) {}

// ExitFormatAnnotation is called when production formatAnnotation is exited.
func (s *BaseSwaggableParserListener) ExitFormatAnnotation(ctx *FormatAnnotationContext) {}

// EnterAnnotationType is called when production annotationType is entered.
func (s *BaseSwaggableParserListener) EnterAnnotationType(ctx *AnnotationTypeContext) {}

// ExitAnnotationType is called when production annotationType is exited.
func (s *BaseSwaggableParserListener) ExitAnnotationType(ctx *AnnotationTypeContext) {}

// EnterParameterDefinition is called when production parameterDefinition is entered.
func (s *BaseSwaggableParserListener) EnterParameterDefinition(ctx *ParameterDefinitionContext) {}

// ExitParameterDefinition is called when production parameterDefinition is exited.
func (s *BaseSwaggableParserListener) ExitParameterDefinition(ctx *ParameterDefinitionContext) {}

// EnterObjElement is called when production objElement is entered.
func (s *BaseSwaggableParserListener) EnterObjElement(ctx *ObjElementContext) {}

// ExitObjElement is called when production objElement is exited.
func (s *BaseSwaggableParserListener) ExitObjElement(ctx *ObjElementContext) {}

// EnterModelElement is called when production modelElement is entered.
func (s *BaseSwaggableParserListener) EnterModelElement(ctx *ModelElementContext) {}

// ExitModelElement is called when production modelElement is exited.
func (s *BaseSwaggableParserListener) ExitModelElement(ctx *ModelElementContext) {}

// EnterModelElements is called when production modelElements is entered.
func (s *BaseSwaggableParserListener) EnterModelElements(ctx *ModelElementsContext) {}

// ExitModelElements is called when production modelElements is exited.
func (s *BaseSwaggableParserListener) ExitModelElements(ctx *ModelElementsContext) {}
