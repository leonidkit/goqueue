package int64arr

import (
	"math/rand"
	"sort"
	"time"
)

// Int64Arr is structure that implements the IArray contract with []int64
type Int64Arr struct {
	data []int64
}

// New return Int64Array object, takes the array to be sorted
func New(arr []int64) *Int64Arr {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	return &Int64Arr{
		data: arr,
	}
}

// Median returns median of the array
func (a *Int64Arr) Median() int {
	return int(a.data[(len(a.data)-1)/2])
}

// Max returns max element of the array
func (a *Int64Arr) Max() int {
	return int(a.data[len(a.data)-1])
}

// Min returns min element of the array
func (a *Int64Arr) Min() int {
	return int(a.data[0])
}

// GenerateRandomSlice generate []int64 slice
func GenerateRandomSlice(size int) []int64 {
	slice := make([]int64, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = int64(rand.Intn(999) - rand.Intn(999))
	}
	return slice
}
