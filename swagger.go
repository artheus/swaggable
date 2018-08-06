package main

type ExtendedComponent struct {
	AllOf []interface{}
}

type Component struct {
	Name       string `yaml:"-"`
	Type       string
	Extends    []string `yaml:"-"`
	Properties map[string]*ComponentProperty
	Required   []string `yaml:",omitempty"`
}

type Enum struct {
	Name     string `yaml:"-"`
	TypeName string `yaml:"type"`
	Values   []string
}

func NewComponent(name string) *Component {
	var comp = new(Component)
	comp.Name = name
	comp.Type = "object"
	comp.Properties = make(map[string]*ComponentProperty)
	return comp
}

func (s *Component) addProperty(prop *ComponentProperty) {
	s.Properties[prop.Name] = prop
	if prop.Required == true {
		s.Required = append(s.Required, prop.Name)
	}
}

// TODO: Implement this for real.
func (s *Component) compileExtends() {
	var exComp = new(ExtendedComponent)
	for _, c := range s.Extends {
		var exType = new(Type)
		exType.Ref = "#/components/schemas/" + c
		exComp.AllOf = append(exComp.AllOf, exType)
	}
	exComp.AllOf = append(exComp.AllOf, s)
}

type Type struct {
	Name  string `yaml:"type,omitempty"`
	Items *Type  `yaml:",omitempty"`
	Ref   string `yaml:"$ref,omitempty"`
}

func (s *Type) GetText() {

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
