package Week_04

//遍历
//发现岛屿，则将相连的岛屿变成0
func numIslands(grid [][]byte) int {
	lands := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				lands++
				removeLand(&grid, i, j)
			}
		}
	}
	return lands
}

func removeLand(grid *[][]byte, i, j int) {
	(*grid)[i][j] = '0'
	if j > 0 && (*grid)[i][j-1] == '1' {
		removeLand(grid, i, j-1)
	}
	if j < len((*grid)[i])-1 && (*grid)[i][j+1] == '1' {
		removeLand(grid, i, j+1)
	}
	if i > 0 && (*grid)[i-1][j] == '1' {
		removeLand(grid, i-1, j)
	}
	if i < len(*grid)-1 && (*grid)[i+1][j] == '1' {
		removeLand(grid, i+1, j)
	}
}
