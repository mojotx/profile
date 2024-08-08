package fnord

import (
	"testing"
)

func TestGetRandomRune(t *testing.T) {
	r := GetRandomRune()
	if r < 'A' || r > 'Z' {
		t.Errorf("Expected rune between 'A' and 'Z', got %q", r)
	}
}
func TestGetRandomString(t *testing.T) {
	length := 10
	s := GetRandomString(length)
	if len(s) != length {
		t.Errorf("Expected string of length %d, got %d", length, len(s))
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	s := GetWorkingData()
	for i := 0; i < b.N; i++ {
		BubbleSort(s)
	}
}

type TestData[T SortableConstraint] struct {
	name  string
	slice []T
	want  bool
}

func TestSliceIsSorted(t *testing.T) {

	var testData = []TestData{
		{
			name:  "int unsorted slice",
			slice: []int{8, 6, 7, 5, 3, 0, 9},
			want:  false,
		},
		{
			name:  "int sorted slice",
			slice: []int{1, 2, 3, 5, 8, 13, 21},
			want:  true,
		},
		{
			name:  "float64 unsorted slice",
			slice: []float64{8.0, 6.0, 7.0, 5.0, 3.0, 0.0, 9.0},
			want:  false,
		},
		{
			name:  "float64 sorted slice",
			slice: []float64{1.0, 2.0, 3.0, 5.0, 8.0, 13.0, 21.0},
			want:  true,
		},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			got := SliceIsSorted(td.slice)
			if got != td.want {
				t.Errorf("Expected %v, got %v", td.want, got)
			}
		})
	}
}
