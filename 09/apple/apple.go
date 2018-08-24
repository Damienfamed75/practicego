package apple

// Apple is a fruit used to test with
type Apple struct {
	Color string // string are a collection of characters in a slice or array

	hasWorm bool // bool (booleans) have a true or false value
}

// NewApple creates a new apple for the examiner to use
func NewApple(col string, worm bool) *Apple {

	return &Apple{
		Color:   col,
		hasWorm: worm,
	}
}

// Worm gets the value of hasWorm
func (a *Apple) Worm() bool {
	return a.hasWorm
}
