// Code generated from SwaggableParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // SwaggableParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SwaggableParserListener is a complete listener for a parse tree produced by SwaggableParser.
type SwaggableParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterObjDeclaration is called when entering the objDeclaration production.
	EnterObjDeclaration(c *ObjDeclarationContext)

	// EnterEnumDeclaration is called when entering the enumDeclaration production.
	EnterEnumDeclaration(c *EnumDeclarationContext)

	// EnterEnumBlock is called when entering the enumBlock production.
	EnterEnumBlock(c *EnumBlockContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterExtendsStatement is called when entering the extendsStatement production.
	EnterExtendsStatement(c *ExtendsStatementContext)

	// EnterObjTail is called when entering the objTail production.
	EnterObjTail(c *ObjTailContext)

	// EnterNativeType is called when entering the nativeType production.
	EnterNativeType(c *NativeTypeContext)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterMapType is called when entering the mapType production.
	EnterMapType(c *MapTypeContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterFormatIdentifier is called when entering the formatIdentifier production.
	EnterFormatIdentifier(c *FormatIdentifierContext)

	// EnterFormatType is called when entering the formatType production.
	EnterFormatType(c *FormatTypeContext)

	// EnterFormatAnnotation is called when entering the formatAnnotation production.
	EnterFormatAnnotation(c *FormatAnnotationContext)

	// EnterAnnotationType is called when entering the annotationType production.
	EnterAnnotationType(c *AnnotationTypeContext)

	// EnterParameterDefinition is called when entering the parameterDefinition production.
	EnterParameterDefinition(c *ParameterDefinitionContext)

	// EnterObjElement is called when entering the objElement production.
	EnterObjElement(c *ObjElementContext)

	// EnterModelElement is called when entering the modelElement production.
	EnterModelElement(c *ModelElementContext)

	// EnterModelElements is called when entering the modelElements production.
	EnterModelElements(c *ModelElementsContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitObjDeclaration is called when exiting the objDeclaration production.
	ExitObjDeclaration(c *ObjDeclarationContext)

	// ExitEnumDeclaration is called when exiting the enumDeclaration production.
	ExitEnumDeclaration(c *EnumDeclarationContext)

	// ExitEnumBlock is called when exiting the enumBlock production.
	ExitEnumBlock(c *EnumBlockContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitExtendsStatement is called when exiting the extendsStatement production.
	ExitExtendsStatement(c *ExtendsStatementContext)

	// ExitObjTail is called when exiting the objTail production.
	ExitObjTail(c *ObjTailContext)

	// ExitNativeType is called when exiting the nativeType production.
	ExitNativeType(c *NativeTypeContext)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitMapType is called when exiting the mapType production.
	ExitMapType(c *MapTypeContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitFormatIdentifier is called when exiting the formatIdentifier production.
	ExitFormatIdentifier(c *FormatIdentifierContext)

	// ExitFormatType is called when exiting the formatType production.
	ExitFormatType(c *FormatTypeContext)

	// ExitFormatAnnotation is called when exiting the formatAnnotation production.
	ExitFormatAnnotation(c *FormatAnnotationContext)

	// ExitAnnotationType is called when exiting the annotationType production.
	ExitAnnotationType(c *AnnotationTypeContext)

	// ExitParameterDefinition is called when exiting the parameterDefinition production.
	ExitParameterDefinition(c *ParameterDefinitionContext)

	// ExitObjElement is called when exiting the objElement production.
	ExitObjElement(c *ObjElementContext)

	// ExitModelElement is called when exiting the modelElement production.
	ExitModelElement(c *ModelElementContext)

	// ExitModelElements is called when exiting the modelElements production.
	ExitModelElements(c *ModelElementsContext)
}
