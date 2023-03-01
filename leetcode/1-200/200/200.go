/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/2/27 下午5:37.
 * Author:  guojia(https://github.com/guojia99)
 */

package _00

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
//岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
//此外，你可以假设该网格的四条边均被水包围。
//
//
//
//示例 1：
//
//输入：grid = [
//["1","1","1","1","0"],
//["1","1","0","1","0"],
//["1","1","0","0","0"],
//["0","0","0","0","0"]
//]
//输出：1
//示例 2：
//
//输入：grid = [
//["1","1","0","0","0"],
//["1","1","0","0","0"],
//["0","0","1","0","0"],
//["0","0","0","1","1"]
//]
//输出：3
//
//
//提示：
//
//m == grid.length
//n == grid[i].length
//1 <= m, n <= 300
//grid[i][j] 的值为 '0' 或 '1'

func numIslands(grid [][]byte) int {
	ans := 0

	var run func(i, j int)
	run = func(i, j int) {
		if i >= len(grid) || j >= len(grid[0]) || i < 0 || j < 0 {
			return
		}

		if grid[i][j] != '1' {
			return
		}

		grid[i][j] = '2'
		run(i+1, j)
		run(i-1, j)
		run(i, j+1)
		run(i, j-1)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ans += 1
				run(i, j)
			}

		}
	}
	return ans
}
