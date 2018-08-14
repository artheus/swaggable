package swagger

type Root struct {
	Paths      map[string]*Path   `yaml:"paths,omitempty"`
	Components *ComponentsWrapper `yaml:"components,omitempty"`
}

func NewRoot() *Root {
	r := new(Root)
	r.Components = NewComponentsWrapper()
	return r
}

func (s *Root) AddPath(name string, p *Path) {
	if s.Paths == nil {
		s.Paths = make(map[string]*Path)
	}

	s.Paths[name] = p
}

type ComponentsWrapper struct {
	Schemas    map[string]*Component `yaml:"schemas,omitempty"`
	Parameters map[string]*Parameter `yaml:"parameters,omitempty"`
}

func NewComponentsWrapper() *ComponentsWrapper {
	cw := new(ComponentsWrapper)
	cw.Schemas = make(map[string]*Component)
	cw.Parameters = make(map[string]*Parameter)
	return cw
}

func (s *ComponentsWrapper) AddSchema(name string, c *Component) {
	if s.Schemas == nil {
		s.Schemas = make(map[string]*Component)
	}

	s.Schemas[name] = c
}

func (s *ComponentsWrapper) AddSchemas(schemaMap map[string]*Component) {
	for name, schema := range schemaMap {
		s.AddSchema(name, schema)
	}
}

type Path struct {
	Get     *Operation `yaml:"get,omitempty"`
	Post    *Operation `yaml:"post,omitempty"`
	Put     *Operation `yaml:"put,omitempty"`
	Patch   *Operation `yaml:"patch,omitempty"`
	Delete  *Operation `yaml:"delete,omitempty"`
	Head    *Operation `yaml:"head,omitempty"`
	Options *Operation `yaml:"options,omitempty"`
	Trace   *Operation `yaml:"trace,omitempty"`
}

func NewPath() *Path {
	p := new(Path)
	return p
}

type Operation struct {
	Tags        []string             `yaml:"tags,omitempty"`
	OperationId string               `yaml:"operationId,omitempty"`
	Parameters  []*Parameter         `yaml:"parameters,omitempty"`
	Responses   map[string]*Response `yaml:"responses,omitempty"`
	Schema      string               `yaml:"schema,omitempty"`
	Ref         string               `yaml:"$ref,omitempty"`
	Description string               `yaml:"description,omitempty"`
}

func NewOperation() *Operation {
	o := new(Operation)
	return o
}

func (s *Operation) AddTag(tag string) {
	if s.Tags == nil {
		s.Tags = []string{}
	}

	s.Tags = append(s.Tags, tag)
}

func (s *Operation) AddParameter(param *Parameter) {
	if s.Parameters == nil {
		s.Parameters = []*Parameter{}
	}

	s.Parameters = append(s.Parameters, param)
}

func (s *Operation) AddResponse(name string, response *Response) {
	if s.Responses == nil {
		s.Responses = make(map[string]*Response)
	}

	s.Responses[name] = response
}

type Parameter struct {
	Name     string     `yaml:"name,omitempty"`
	In       string     `yaml:"in,omitempty"`
	Required bool       `yaml:"required"`
	Schema   *Component `yaml:"schema"`
}

func NewParameter() *Parameter {
	p := new(Parameter)
	p.Schema = &Component{}
	return p
}

type Response struct {
	Content     map[string]*ResponseContent `yaml:"content,omitempty"`
	Description string                      `yaml:"description,omitempty"`
}

func NewResponse() *Response {
	r := new(Response)
	return r
}

func (s *Response) AddContent(name string, content *ResponseContent) {
	if s.Content == nil {
		s.Content = make(map[string]*ResponseContent)
	}

	s.Content[name] = content
}

type ResponseContent struct {
	Schema *Component
}

func NewResponseContent() *ResponseContent {
	rc := new(ResponseContent)
	rc.Schema = &Component{}
	return rc
}
