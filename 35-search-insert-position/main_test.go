package search_insert_position

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	nums   []int
	target int

	want int
}

var twoSumTests = []testcase{
	{[]int{1, 3, 5, 6}, 5, 2},
	{[]int{1, 3, 5, 6}, 2, 1},
	{[]int{1, 3, 5, 6}, 7, 4},
}

func TestCases(t *testing.T) {
	for i, tc := range twoSumTests {
		t.Run(fmt.Sprintf("tc-%d", i+1), func(t *testing.T) {
			output := searchInsert(tc.nums, tc.target)
			require.Equal(t, tc.want, output)
		})
	}
}
