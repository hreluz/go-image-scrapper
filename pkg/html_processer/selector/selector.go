package selector

type SelectorName string
type SelectorType string
type SelectorTypes []SelectorType

const (
	ID    SelectorType = "id"
	CLASS SelectorType = "class"
)

var SELECTOR_TYPE_OPTIONS = SelectorTypes{ID, CLASS}

type Selector struct {
	stype SelectorType
	name  SelectorName
}

// New returns a new Selector
func New(stype SelectorType, name string) *Selector {
	return &Selector{
		stype: stype,
		name:  SelectorName(name),
	}
}

func (s *Selector) GetType() SelectorType {
	return s.stype
}

func (s *Selector) GetName() SelectorName {
	return s.name
}
