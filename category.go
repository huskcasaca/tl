package tl

// Category of Definition.
type Category byte

const (
	CategoryUnknown Category = iota
	CategoryPredict
	CategoryFunction
)

func (c Category) String() string {
	switch c {
	case CategoryPredict:
		return "predict"
	case CategoryFunction:
		return "function"
	default:
		return "<UNKNOWN>"
	}
}
