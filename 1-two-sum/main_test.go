package two_sum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	nums   []int
	target int

	want []int
}

var twoSumTests = []testcase{
	{nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
	{nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
	{nums: []int{3, 3}, target: 6, want: []int{0, 1}},
}

func TestTwoSum(t *testing.T) {
	for _, tc := range twoSumTests {
		output := twoSum(tc.nums, tc.target)
		require.Equal(t, tc.want, output)
	}
}
