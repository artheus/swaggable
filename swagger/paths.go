package swagger

import (
	"strings"

	"github.com/iancoleman/strcase"
)

func CreateGetAndDeleteByIndexPath(comp *Component, index *ComponentProperty) (string, *Path) {
	path := NewPath()

	path.Delete = NewOperation()
	path.Get = NewOperation()

	path.Delete.Description = "Delete one " + strcase.ToCamel(comp.Name)
	path.Get.Description = "Get one existing " + strcase.ToCamel(comp.Name)

	path.Delete.OperationId = "delete" + comp.Name + "By" + strcase.ToCamel(index.Name)
	path.Get.OperationId = "get" + comp.Name + "By" + strcase.ToCamel(index.Name)

	pathIdParam := NewParameter()
	pathIdParam.In = "path"
	pathIdParam.Name = strings.ToLower(index.Name)
	pathIdParam.Required = true
	pathIdParam.Schema.Ref = index.PropRef
	pathIdParam.Schema.Type = index.PropType
	path.Get.AddParameter(pathIdParam)
	path.Delete.AddParameter(pathIdParam)

	getOkResponse := NewResponse()
	getOkResponse.Description = "OK"
	responseContent := NewResponseContent()
	responseContent.Schema.Ref = "#/components/schemas/" + comp.Name
	getOkResponse.AddContent("application/json", responseContent)
	path.Get.AddResponse("200", getOkResponse)

	deleteOkResponse := NewResponse()
	deleteOkResponse.Description = "OK"
	path.Delete.AddResponse("200", deleteOkResponse)

	pathNotFoundResponse := NewResponse()
	pathNotFoundResponse.Description = "NOT FOUND"
	path.Get.AddResponse("404", pathNotFoundResponse)
	path.Delete.AddResponse("404", pathNotFoundResponse)

	internalErrorResponse := NewResponse()
	internalErrorResponse.Description = "INTERNAL SERVER ERROR"
	path.Get.AddResponse("500", internalErrorResponse)
	path.Delete.AddResponse("500", internalErrorResponse)

	return "/" + strings.ToLower(comp.Name) + "/{" + strings.ToLower(index.Name) + "}", path
}

func CreateUpdateAndCreatePath(comp *Component) (string, *Path) {
	path := NewPath()

	path.Put = NewOperation()
	path.Put.OperationId = "update" + comp.Name

	path.Post = NewOperation()
	path.Post.OperationId = "create" + comp.Name

	path.Put.Description = "Update existing instance of " + strcase.ToCamel(comp.Name)
	path.Post.Description = "Create one new " + strcase.ToCamel(comp.Name)

	pathComponentParam := NewParameter()
	pathComponentParam.Name = strcase.ToLowerCamel(comp.Name)
	pathComponentParam.In = "query"
	pathComponentParam.Required = true
	pathComponentParam.Schema.Ref = "#/components/schemas/" + comp.Name

	path.Put.AddParameter(pathComponentParam)
	path.Post.AddParameter(pathComponentParam)

	pathOkResponse := NewResponse()
	pathOkResponse.Description = "OK"

	responseContent := NewResponseContent()
	responseContent.Schema.Ref = "#/components/schemas/" + comp.Name

	pathOkResponse.AddContent("application/json", responseContent)

	path.Put.AddResponse("200", pathOkResponse)
	path.Post.AddResponse("200", pathOkResponse)

	pathNotFoundResponse := NewResponse()
	pathNotFoundResponse.Description = "NOT FOUND"
	path.Put.AddResponse("404", pathNotFoundResponse)

	internalErrorResponse := NewResponse()
	internalErrorResponse.Description = "INTERNAL SERVER ERROR"
	path.Put.AddResponse("500", internalErrorResponse)
	path.Post.AddResponse("500", internalErrorResponse)

	return "/" + strings.ToLower(comp.Name), path
}
