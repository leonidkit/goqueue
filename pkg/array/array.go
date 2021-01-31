package array

// IArray defines a array contract
type IArray interface {
	// Returns the median of array
	Median() int

	// Returns the max element of array
	Max() int

	// Returns the min element of array
	Min() int
}
