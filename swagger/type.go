package swagger

type TypeInterface interface {
	GetText() string
}

type Type struct {
	Name  string `yaml:"type,omitempty"`
	Items *Type  `yaml:",omitempty"`
	Ref   string `yaml:"$ref,omitempty"`
}

func NewType() *Type {
	t := new(Type)
	return t
}

func (s *Type) GetText() string {
	if s.Ref != "" {
		return "#/components/schemas/" + s.Ref
	}

	return s.Name
}
