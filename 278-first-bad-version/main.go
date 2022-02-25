package first_bad_version

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	return version < 5
}

func firstBadVersion(n int) int {
	a := 0
	b := n

	var (
		mid int
		ibv bool
	)
	for {
		mid = a + (b-a)/2

		ibv = isBadVersion(mid)

		if ibv && b-a <= 1 {
			return mid
		}

		if ibv {
			b = mid

			continue
		}

		a = mid + 1
	}
}
