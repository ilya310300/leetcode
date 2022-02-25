package binary_search

func search(nums []int, target int) int {
	a := 0
	b := len(nums)

	var mid int
	for {
		if len(nums[a:b]) == 0 {
			return -1
		}

		mid = a + (b-a)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			b = mid

			continue
		}

		if nums[mid] < target {
			a = mid + 1

			continue
		}
	}

	return -1
}
