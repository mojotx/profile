package fnord

import (
	"crypto/rand"
	"math/big"
	"strings"

	"github.com/fatih/color"
)

const (
	// Start with 1000
	numberOfWorkingStrings = 100
)

func GetWorkingData() (data []string) {
	data = make([]string, numberOfWorkingStrings)
	// Create a massive slice of random strings
	for i := 0; i < numberOfWorkingStrings; i++ {
		data[i] = GetRandomString(120)
	}
	return data
}

func BubbleSort(data []string) {
	n := len(data)
	var ctr uint64
	for i := 0; i < n; i++ {
		ctr++
		for j := 0; j < n-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	color.HiBlack("BubbleSort took %d iterations\n", ctr)
}

// QuickSort sorts a slice of strings in ascending order using the quick sort algorithm.
func QuickSort(data []string) {
	if len(data) <= 1 {
		return // Base case: nothing to sort
	}

	pivotIndex := partition(data)
	QuickSort(data[:pivotIndex])   // Recursively sort elements less than pivot
	QuickSort(data[pivotIndex+1:]) // Recursively sort elements greater than pivot
}

// partition rearranges the elements around the pivot (last element) and returns the index of the pivot.
func partition(data []string) int {
	n := len(data)
	pivot := data[n-1]
	i := 0

	for j := 0; j < n-1; j++ {
		if data[j] <= pivot {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}

	// Move pivot to its final sorted position
	data[i], data[len(data)-1] = data[len(data)-1], data[i]
	return i
}

func GetRandomRune() rune {
	max := big.NewInt(25)
	nBig, err := rand.Int(rand.Reader, max.Add(max, big.NewInt(1)))
	if err != nil {
		panic(err.Error())
	}

	n := nBig.Int64()

	return rune(n + 65)
}

func GetRandomString(length int) string {
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(GetRandomRune())
	}
	return b.String()
}

type SortableConstraint interface {
	//constraints.Integer | constraints.Float | constraints.String
	int | uint64 | float64 | string
	// LessThan(other interface{}) bool
}

/*
func (i int) LessThan(other interface{}) bool {
	j, ok := other.(int)
	if !ok {
		return false
	}
	return i < j
}
*/

func SliceIsSorted[T SortableConstraint](slice []T) bool {
	for i := 0; i < len(slice); i++ {
		switch {
		case i == 0:
			return true
		case i > 0 && slice[i] < slice[i-1]:
			return false
		}
	}
	return true
}
