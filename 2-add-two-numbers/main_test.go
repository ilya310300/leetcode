package add_two_numbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	l1 *ListNode
	l2 *ListNode

	want *ListNode
}

var twoSumTests = []testcase{
	{
		l1: &ListNode{
			Val: 1,
		},
		l2: &ListNode{
			Val: 2,
		},
		want: &ListNode{
			Val: 3,
		},
	},
	{
		l1: &ListNode{
			Val: 5,
		},
		l2: &ListNode{
			Val: 9,
		},
		want: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 1,
			},
		},
	},
	{
		l1: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
		l2: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
		want: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 8,
				},
			},
		},
	},
}

func TestTwoSum(t *testing.T) {
	for i, tc := range twoSumTests {
		t.Run(fmt.Sprintf("tc-%d", i+1), func(t *testing.T) {
			output := addTwoNumbers(tc.l1, tc.l2)
			require.Equal(t, tc.want, output)
		})
	}
}
