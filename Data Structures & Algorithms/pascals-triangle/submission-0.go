func generate(numRows int) [][]int {
    dp := make([][]int, numRows)
    result := make([][]int, numRows)
    for i := 0; i < numRows; i++ {
        dp[i] = make([]int, numRows)
        for j := 0; j <= i; j++ {
            if i - 1 >= 0 && j - 1 >= 0 {
                dp[i][j] += dp[i - 1][j - 1]
            }
            if i - 1 >= 0 && j >= 0 {
                dp[i][j] += dp[i - 1][j]
            }
            if i == 0 && j == 0 {
                dp[i][j] = 1
            } 
            if dp[i][j] != 0 { result[i] = append(result[i], dp[i][j])}
        }
    }

    return result
    
}
