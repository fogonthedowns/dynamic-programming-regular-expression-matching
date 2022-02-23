package main

// Pair is to be used in memo below as a dynamic programming technique
type Pair struct {
	i int
	j int
}

func isMatch(s string, p string) bool {
	var memo = map[Pair]bool{} // declare this globally will cause problems in Leetcode
	return dp(0, 0, s, p, memo)
}

func dp(i, j int, s, p string, memo map[Pair]bool) bool {
	var ans bool

	// if (i, j) not in memo
	if _, ok := memo[Pair{i, j}]; !ok {
		if j == len(p) {
			// terminating condition
			ans = i == len(s)
		} else {
			// main logic
			firstMatch := i < len(s) && (p[j] == s[i] || p[j] == '.')
			if j < len(p)-1 && p[j+1] == '*' {
				zeroOcr := dp(i, j+2, s, p, memo)
				if firstMatch {
					repeat := dp(i+1, j, s, p, memo)
					ans = zeroOcr || repeat
				} else {
					ans = zeroOcr
				}
			} else {
				ans = firstMatch && dp(i+1, j+1, s, p, memo)
			}
			memo[Pair{i, j}] = ans
		}
	} else {
		// if in memo / !ok
		ans = memo[Pair{i, j}]
	}

	return ans
}
