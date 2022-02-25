package search_insert_position

func searchInsert(nums []int, target int) int {
	a := 0
	b := len(nums)

	var mid int
	for {
		if len(nums[a:b]) == 0 {
			return a
		}

		mid = a + (b-a)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			b = mid

			continuew
		}

		if nums[mid] < target {
			a = mid + 1

			continue
		}
	}
}
