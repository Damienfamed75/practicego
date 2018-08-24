package calculator

// Calculator carries various mathematical functions to ease your pain.
type Calculator struct{}

// New returns a pointer to a calculator.
func New() *Calculator {
	return &Calculator{}
}

// Add takes in a variadic number of integers and returns the sum.
func (c *Calculator) Add(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
