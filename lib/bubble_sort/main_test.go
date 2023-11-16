package bubble_sort_test

import (
	"reflect"
	"testing"

	"github.com/cherryramatisdev/algorithms-apocalipse/lib/bubble_sort"
)

func TestHappy(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}

	bubble_sort.BubbleSort(arr)
	want := []int{3, 4, 7, 9, 42, 69, 420}

	if !reflect.DeepEqual(arr, want) {
		t.Errorf("Expected %v, got %v", want, arr)
	}
}
