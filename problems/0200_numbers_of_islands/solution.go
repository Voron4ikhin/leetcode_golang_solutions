package _200_numbers_of_islands

func NumIslands(grid [][]byte) int {
	var result int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				dfs(i, j, grid)
				result++
			}
		}
	}

	return result

}

func dfs(i, j int, grid [][]byte) {
	m := len(grid)
	n := len(grid[0])

	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
		return
	}

	grid[i][j] = '2'

	dfs(i-1, j, grid)
	dfs(i+1, j, grid)
	dfs(i, j-1, grid)
	dfs(i, j+1, grid)
}
