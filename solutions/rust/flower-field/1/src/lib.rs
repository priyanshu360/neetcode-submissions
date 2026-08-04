pub fn annotate(garden: &[&str]) -> Vec<String> {
    let mut res = Vec::new();

    let dir: [i32; 3] = [1, -1, 0];
    for i in 0..garden.len() {
        let mut value = String::new();
        for (j, c) in garden[i].as_bytes().iter().enumerate() {
            if *c == '*' as u8 {
                value += "*";
                continue;
            }

            let mut count = 0;
            for dx in dir {
                for dy in dir {
                    let nx = i as i32 + dx;
                    let ny = j as i32 + dy;

                    if nx < 0 || nx >= garden.len() as i32 {
                        continue;
                    }
                    if ny < 0 || ny >= garden[0].len() as i32 {
                        continue;
                    }

                    count += if garden[nx as usize].as_bytes()[ny as usize] == b'*' {
                        1
                    } else {
                        0
                    };
                }
            }

            if count == 0 {
                value.push(' ');
            } else {
                value += &count.to_string();
            }
        }
        res.push(value);
    }

    return res;
}
