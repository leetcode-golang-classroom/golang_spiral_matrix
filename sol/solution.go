package sol

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	top, bottom := 0, len(matrix)
	left, right := 0, len(matrix[0])
	for top < bottom && left < right {
		//top layer: go left
		for col := left; col < right; col++ {
			result = append(result, matrix[top][col])
		}
		top++
		//right layer: go down
		for row := top; row < bottom; row++ {
			result = append(result, matrix[row][right-1])
		}
		right--
		if top == bottom || left == right {
			break
		}
		//bottom layer: go right
		for col := right - 1; col >= left; col-- {
			result = append(result, matrix[bottom-1][col])
		}
		bottom--
		//left layer: go up
		for row := bottom - 1; row >= top; row-- {
			result = append(result, matrix[row][left])
		}
		left++
	}
	return result
}
