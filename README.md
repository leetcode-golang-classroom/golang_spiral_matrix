# golang_spiral_matrix

Given an `m x n` `matrix`, return *all elements of the* `matrix` *in spiral order*.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2020/11/13/spiral1.jpg](https://assets.leetcode.com/uploads/2020/11/13/spiral1.jpg)

```
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

```

**Example 2:**

![https://assets.leetcode.com/uploads/2020/11/13/spiral.jpg](https://assets.leetcode.com/uploads/2020/11/13/spiral.jpg)

```
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]

```

**Constraints:**

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 10`
- `100 <= matrix[i][j] <= 100`

## 解析

給定一個矩陣 matrix 

要求寫一個演算法透過 順時針螺旋的順序來把矩陣轉換成陣列

 

可以注意到 spiral order 

每次拆分成成四個部份 

最上層由左往右走訪

最右層由上往下走訪

最下層由右往左走訪

最左層由下往上走訪


![](https://i.imgur.com/FDJAwq7.png)
可以注意到每次走訪完一層 就可以把走訪的該層 排除，範圍縮小

比如說當最上層走到最右邊時，就可以把 top 更新為 top +=1

這樣每次移動 top, left , bottom, right 指標直到 top == bottom || left == right

逐步把每個走訪的元素放到新的陣列內

![](https://i.imgur.com/YVmjSd7.png)

時間複雜度是 O(m*n) where n = len(matrix), m = len(matrix[0])

空間複雜度是 O(m*n) 回傳的陣列長度

## 程式碼
```go
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

```
## 困難點

1. 需要思考出如何有系統的去走訪整個矩陣
2. 需要想出更新四個走訪範圍的方式

## Solve Point

- [x]  初始化 left = 0, top = 0, bottom = len(matrix), right = len(matrix[0]), result = []
- [x]  當 left < right 且 top < bottom 時 做以下走訪
- [x]  i := left.. right -1 ,新增 matrix[top][i] 到 result
- [x]  更新 top+= 1
- [x]  i := top-1.. bottom-1 ,新增 matrix[i][right-1] 到 result
- [x]  更新 right -= 1
- [x]  檢查 如果 left == right 或是 top == bottom 就停止回圈
- [x]  i := right -1.. left ,新增 matrix[bottom-1][i] 到 result
- [x]  更新 bottom -= 1
- [x]  i := bottom -1.. top ,新增 matrix[i][left] 到 result
- [x]  更新 left += 1
- [x]  回傳 result