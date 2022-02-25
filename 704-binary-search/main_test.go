package binary_search

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
	{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},  // 9 exists in nums and its index is 4
	{[]int{-1, 0, 3, 5, 9, 12}, 2, -1}, // 2 does not exist in nums so return -1
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for i, tc := range twoSumTests {
		t.Run(fmt.Sprintf("tc-%d", i+1), func(t *testing.T) {
			output := search(tc.nums, tc.target)
			require.Equal(t, tc.want, output)
		})
	}
}
