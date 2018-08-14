package mapping

import (
	"strings"

	"github.com/artheus/swaggable/antlr/parser"

	"github.com/artheus/swaggable/swagger"
)

type BaseElement struct {
	*swagger.Component
}

func NewBaseElement(name string) *BaseElement {
	be := new(BaseElement)
	be.Component = swagger.NewComponent(name)
	return be
}

func CreateComponent(ctx *parser.ObjDeclarationContext) *swagger.Component {
	var comp = swagger.NewComponent(ctx.Identifier().GetText())
	var objTail = ctx.ObjTail().(*parser.ObjTailContext)

	if objTail.ExtendsStatement() != nil {
		decorateExtends(comp, objTail.ExtendsStatement().(*parser.ExtendsStatementContext))
	}

	for _, objElm := range objTail.AllObjElement() {
		decorateParameter(comp, objElm.(*parser.ObjElementContext).ParameterDefinition().(*parser.ParameterDefinitionContext))
	}

	return comp
}

func CreateEnum(ctx *parser.EnumDeclarationContext) *swagger.Component {
	var enum = swagger.NewComponent(ctx.Identifier().GetText())

	t := typeContextToType(ctx.TypeName().(*parser.TypeNameContext))
	enum.Type = t.GetText()

	for _, enumElm := range ctx.EnumBlock().(*parser.EnumBlockContext).AllIdentifier() {
		enum.EnumValues = append(enum.EnumValues, enumElm.GetText())
	}

	return enum
}

func CreateBase(ctx *parser.BaseDeclarationContext) *BaseElement {
	base := NewBaseElement(ctx.Identifier().GetText())

	var objTail = ctx.ObjTail().(*parser.ObjTailContext)

	if objTail.ExtendsStatement() != nil {
		decorateExtends(base.Component, objTail.ExtendsStatement().(*parser.ExtendsStatementContext))
	}

	for _, objElm := range objTail.AllObjElement() {
		decorateParameter(base.Component, objElm.(*parser.ObjElementContext).ParameterDefinition().(*parser.ParameterDefinitionContext))
	}

	return base
}

func decorateExtends(comp *swagger.Component, ctx *parser.ExtendsStatementContext) {
	// Collect extends names
	for _, tkn := range ctx.IdentifierList().(*parser.IdentifierListContext).AllIdentifier() {
		comp.Extends = append(comp.Extends, tkn.GetText())
	}
}

func decorateParameter(comp *swagger.Component, ctx *parser.ParameterDefinitionContext) {
	var property = swagger.NewComponentProperty()

	property.Name = ctx.Identifier().GetText()
	decoratePropertyType(property, ctx.TypeName().(*parser.TypeNameContext))

	for _, anType := range ctx.AllAnnotationType() {
		var actx *parser.AnnotationTypeContext = anType.(*parser.AnnotationTypeContext)
		if actx.Required() != nil {
			property.Required = true
		}

		if actx.Indexed() != nil {
			property.Indexed = true
			property.Searchable = true
		}

		if actx.Searchable() != nil {
			property.Searchable = true
		}

		if actx.FormatAnnotation() != nil {
			property.Format = strings.Replace(actx.FormatAnnotation().(*parser.FormatAnnotationContext).GetText(), "format.", "", 1)
		}
	}

	comp.AddProperty(property)
}

func decoratePropertyType(prop *swagger.ComponentProperty, ctx *parser.TypeNameContext) {
	var t = typeContextToType(ctx)

	if t.Items != nil {
		prop.PropItems = t.Items
	}

	if t.Name != "" {
		prop.PropType = t.Name
	}

	if t.Ref != "" {
		prop.PropRef = t.Ref
	}
}

func typeContextToType(ctx *parser.TypeNameContext) *swagger.Type {
	var t = swagger.NewType()

	if ctx.NativeType() != nil {
		if ctx.NativeType().(*parser.NativeTypeContext).ArrayType() != nil {
			t.Name = "array"
			t.Items = typeContextToType(ctx.NativeType().(*parser.NativeTypeContext).ArrayType().(*parser.ArrayTypeContext).TypeName().(*parser.TypeNameContext))
		} else {
			t.Name = ctx.GetText()
		}
	}

	if ctx.Identifier() != nil {
		t.Ref = "#/components/schemas/" + ctx.GetText()
	}

	return t
}
