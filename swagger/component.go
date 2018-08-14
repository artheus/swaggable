package swagger

type ComponentInterface interface {
	AddProperty(prop *ComponentProperty)
	CompileExtends() *Component
	GetExtends() []string
	GetProperties() map[string]*ComponentProperty
	DeleteExtends(i int)
}

type Component struct {
	Name       string                        `yaml:"-"`
	Type       string                        `yaml:"type,omitempty"`
	Extends    []string                      `yaml:"-"`
	Properties map[string]*ComponentProperty `yaml:"properties,omitempty"`
	Required   []string                      `yaml:",omitempty"`
	AllOf      []*Component                  `yaml:"allOf,omitempty"`
	Ref        string                        `yaml:"$ref,omitempty"`
	EnumValues []string                      `yaml:"enum,omitempty"`
}

type ComponentProperty struct {
	Name       string `yaml:"-"`
	PropType   string `yaml:"type,omitempty"`
	PropRef    string `yaml:"$ref,omitempty"`
	PropItems  *Type  `yaml:"items,omitempty"`
	Required   bool   `yaml:"-"`
	Format     string `yaml:"-"`
	Indexed    bool   `yaml:"-"`
	Searchable bool   `yaml:"-"`
}

func NewComponent(name string) *Component {
	var comp = new(Component)
	comp.Name = name
	comp.Type = "object"
	comp.Properties = make(map[string]*ComponentProperty)
	return comp
}

func NewComponentProperty() *ComponentProperty {
	cp := new(ComponentProperty)
	return cp
}

func (s *Component) AddProperty(prop *ComponentProperty) {
	s.Properties[prop.Name] = prop
	if prop.Required == true {
		s.Required = append(s.Required, prop.Name)
	}
}

func (s *Component) CompileExtends() *Component {
	if len(s.Extends) == 0 {
		return nil
	}

	var exComp = new(Component)
	for _, c := range s.Extends {
		var exType = new(Component)
		exType.Ref = "#/components/schemas/" + c
		exComp.AllOf = append(exComp.AllOf, exType)
	}
	exComp.AllOf = append(exComp.AllOf, s)

	return exComp
}

func (s *Component) GetExtends() []string {
	return s.Extends
}

func (s *Component) DeleteExtends(i int) {
	copy(s.Extends[i:], s.Extends[i+1:])
	s.Extends[len(s.Extends)-1] = "" // or the zero value of T
	s.Extends = s.Extends[:len(s.Extends)-1]
}

func (s *Component) GetProperties() map[string]*ComponentProperty {
	return s.Properties
}
