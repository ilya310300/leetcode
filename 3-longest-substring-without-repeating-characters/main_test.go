package longest_substring_without_repeating_characters

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	s string

	want int
}

var twoSumTests = []testcase{
	{"abcabcbb", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
	{"dvdf", 3},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for i, tc := range twoSumTests {
		t.Run(fmt.Sprintf("tc-%d", i+1), func(t *testing.T) {
			output := lengthOfLongestSubstring(tc.s)
			require.Equal(t, tc.want, output)
		})
	}
}
