package linear_search_test

import (
	"testing"

	"github.com/cherryramatisdev/algorithms-apocalipse/lib/linear_search"
)

func TestHappy(t *testing.T) {
	expects := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	tests := []struct {
		target int
		want   bool
	}{
		{69, true},
		{1336, false},
		{69420, true},
		{69421, false},
		{1, true},
		{0, false},
	}

	for _, tc := range tests {
		got := linear_search.LinearSearch(expects, tc.target)

		if got != tc.want {
			t.Errorf("Expected %t from %d, got %t", tc.want, tc.target, got)
		}
	}
}
