package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]bool, len(s))
	l := 0
	lMax := 0

	// todo: встретив дубль пробегаться по мапе и удалять все символы, чьи индексы хранятся раньше (собственно переделать мапу)
	for _, c := range s {
		if _, ok := m[c]; !ok {
			l++
			m[c] = true

			continue
		}

		if l > lMax {
			lMax = l
		}

		l = 1
		m = make(map[rune]bool, len(s))
		m[c] = true
	}

	if l > lMax {
		lMax = l
	}

	return lMax
}
