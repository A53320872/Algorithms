func isPalin(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func dp(s string, mem map[string]int) int {
	if isPalin(s) {
		mem[s] = 0
		return 0
	}

	if mem[s] != 0 {
		return mem[s]
	}

	minCutting := 99999999
	for i := 1; i < len(s); i++ {
		if isPalin(s[:i]) {
			rightCutting := dp(s[i:], mem)
			if rightCutting+1 < minCutting {
				minCutting = rightCutting + 1
			}
		}
	}

	mem[s] = minCutting
	return mem[s]
}

func minCut(s string) int {
	mem := map[string]int{}
	return dp(s, mem)
}