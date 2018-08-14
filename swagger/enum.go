package swagger

type EnumInterface interface {
	GetText() string
}

type Enum struct {
	Name       string   `yaml:"-"`
	TypeName   string   `yaml:"type"`
	EnumValues []string `yaml:"enum,omitempty"`
}

func NewEnum() *Enum {
	e := new(Enum)
	return e
}

func (s *Enum) GetText() string {
	return s.Name
}
