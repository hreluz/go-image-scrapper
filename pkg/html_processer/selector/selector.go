package selector

type SelectorName string
type SelectorType string
type SelectorTypes []SelectorType

const (
	ID    SelectorType = "id"
	CLASS SelectorType = "class"
	NONE  SelectorType = "none"
)

var SELECTOR_TYPE_OPTIONS = SelectorTypes{ID, CLASS, NONE}

type Selector struct {
	stype SelectorType
	name  SelectorName
}

type SelectorWrapper struct {
	SType SelectorType `json:"stype"`
	Name  SelectorName `json:"name"`
}

func (s *Selector) GetWrapper() *SelectorWrapper {
	return &SelectorWrapper{
		s.stype,
		s.name,
	}
}

// New returns a new Selector
func New(stype SelectorType, name string) *Selector {
	return &Selector{
		stype: stype,
		name:  SelectorName(name),
	}
}

func Empty() *Selector {
	return &Selector{
		stype: NONE,
		name:  SelectorName(""),
	}
}

func (s *Selector) GetType() SelectorType {
	return s.stype
}

func (s *Selector) GetName() SelectorName {
	return s.name
}
