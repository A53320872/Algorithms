func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxOfArray(nums []int) int {
	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}
	return maxNum
}

func dfs(p [2]int, pathLen int, mem map[[2]int]int, row int, col int, matrix [][]int) int {
	i, j := p[0], p[1]
	if mem[[2]int{i, j}] != 0 {
		return mem[[2]int{i, j}]
	}

	down, right, up, left := 0, 0, 0, 0
	if i+1 < row && matrix[i+1][j] > matrix[i][j] {
		down = dfs([2]int{i + 1, j}, pathLen+1, mem, row, col, matrix)
	}
	if j+1 < col && matrix[i][j+1] > matrix[i][j] {
		right = dfs([2]int{i, j + 1}, pathLen+1, mem, row, col, matrix)
	}
	if i-1 >= 0 && matrix[i-1][j] > matrix[i][j] {
		up = dfs([2]int{i - 1, j}, pathLen+1, mem, row, col, matrix)
	}
	if j-1 >= 0 && matrix[i][j-1] > matrix[i][j] {
		left = dfs([2]int{i, j - 1}, pathLen+1, mem, row, col, matrix)
	}

	res := maxOfArray([]int{down, right, up, left}) + 1
	mem[[2]int{i, j}] = res
	return res
}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	row, col := len(matrix), len(matrix[0])
	zeroNodes := [][2]int{}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			minNeighboor := matrix[i][j]
			if i+1 < row {
				minNeighboor = min(minNeighboor, matrix[i+1][j])
			}
			if j+1 < col {
				minNeighboor = min(minNeighboor, matrix[i][j+1])
			}
			if i-1 >= 0 {
				minNeighboor = min(minNeighboor, matrix[i-1][j])
			}
			if j-1 >= 0 {
				minNeighboor = min(minNeighboor, matrix[i][j-1])
			}

			if minNeighboor == matrix[i][j] {
				zeroNodes = append(zeroNodes, [2]int{i, j})
			}
		}
	}
	maxLen := 0
	mem := map[[2]int]int{}
	for i := 0; i < len(zeroNodes); i++ {
		maxLen = maxOfArray([]int{maxLen, dfs(zeroNodes[i], 1, mem, len(matrix), len(matrix[0]), matrix)})
	}
	return maxLen

}