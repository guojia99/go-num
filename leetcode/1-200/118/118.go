package _18

// 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
//
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	lastRows := generate(numRows - 1)
	lastRow := lastRows[len(lastRows)-1]
	var thisRow []int
	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			thisRow = append(thisRow, 1)
			continue
		}
		thisRow = append(thisRow, lastRow[i-1]+lastRow[i])
	}
	return append(lastRows, thisRow)
}
