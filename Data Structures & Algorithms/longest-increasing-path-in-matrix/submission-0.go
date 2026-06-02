func longestIncreasingPath(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])

	indegree := make([][]int, rows)
	for i := range indegree {
		indegree[i] = make([]int, cols)
	}

	dirs := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	// Build indegrees
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]

				if nr < 0 || nr >= rows ||
					nc < 0 || nc >= cols {
					continue
				}

				if matrix[nr][nc] > matrix[r][c] {
					indegree[nr][nc]++
				}
			}
		}
	}

	type Cell struct {
		r, c int
	}

	queue := []Cell{}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if indegree[r][c] == 0 {
				queue = append(queue, Cell{r, c})
			}
		}
	}

	length := 0

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]

			for _, d := range dirs {
				nr, nc := cur.r+d[0], cur.c+d[1]

				if nr < 0 || nr >= rows ||
					nc < 0 || nc >= cols {
					continue
				}

				if matrix[nr][nc] <= matrix[cur.r][cur.c] {
					continue
				}

				indegree[nr][nc]--

				if indegree[nr][nc] == 0 {
					queue = append(queue, Cell{nr, nc})
				}
			}
		}

		length++
	}

	return length
}
