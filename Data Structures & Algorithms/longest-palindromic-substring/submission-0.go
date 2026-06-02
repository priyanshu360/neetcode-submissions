func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	start, maxLen := 0, 1

	expand := func(left, right int) (int, int) {
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
		}
		// Actual palindrome boundaries after over-expanding
		return left + 1, right - left - 1
	}

	for i := 0; i < n; i++ {
		// Odd-length palindrome
		l1, len1 := expand(i, i)
		if len1 > maxLen {
			start = l1
			maxLen = len1
		}

		// Even-length palindrome
		l2, len2 := expand(i, i+1)
		if len2 > maxLen {
			start = l2
			maxLen = len2
		}
	}

	return s[start : start+maxLen]
}