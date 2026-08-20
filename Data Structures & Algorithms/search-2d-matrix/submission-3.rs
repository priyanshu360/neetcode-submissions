impl Solution {
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        if matrix.len() == 0 {
            return false;
        }
        let mut l = 0;
        let mut r = matrix.len() * matrix[0].len();


        while r - l > 1 {
            let m = (l + r) / 2;
            let cr = m / matrix[0].len();
            let cc = m % matrix[0].len();
            if matrix[cr][cc] <= target {
                l = m;
            }
            else {
                r = m;
            }
            // println!("{} {} {}", cr, cc, matrix[cr][cc]);
        }
        let cr = l / matrix[0].len();
        let cc = l % matrix[0].len();
        matrix[cr][cc] == target 
    }
}
