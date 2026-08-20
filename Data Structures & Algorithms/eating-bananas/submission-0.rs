pub fn can_eat(piles: &Vec<i32>, h: i32, x: i32) -> bool {
        let mut left = h;
        for v in piles {
           left -= (v + x - 1) / x;
        }
        left >= 0
    }

impl Solution {
    


    pub fn min_eating_speed(piles: Vec<i32>, h: i32) -> i32 {
        let mut l = 0;
        let mut r = 1_000_000_000;

        while r - l > 1 {
            let m = (l + r) / 2;

            if can_eat(&piles, h, m) {
                r = m;
            }else {
                l = m;
            }
        }

        r
    }
}
