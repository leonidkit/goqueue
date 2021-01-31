package intarr

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

// IntArray is structure that implements the IArray contract with []int
type IntArray struct {
	data []int
}

// New return IntArray object, takes the array to be sorted
func New(arr []int) *IntArray {
	sort.Ints(arr)
	return &IntArray{
		data: arr,
	}
}

// Median returns median of the array
func (a *IntArray) Median() int {
	return a.data[(len(a.data)-1)/2]
}

// Max returns max element of the array
func (a *IntArray) Max() int {
	return a.data[len(a.data)-1]
}

// Min returns min element of the array
func (a *IntArray) Min() int {
	return a.data[0]
}

// GenerateRandomSlice generate []int slice
func GenerateRandomSlice(size int) []int {
	var slice []int
	if size >= math.MaxInt32 {
		slice = make([]int, math.MaxInt32)
	} else {
		slice = make([]int, size)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(size) - rand.Intn(size)
	}
	return slice
}
