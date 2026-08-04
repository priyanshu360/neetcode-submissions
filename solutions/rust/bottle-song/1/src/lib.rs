fn bottles(n: u32) -> String {
    match n {
        0 => "no green bottles".to_string(),
        1 => "one green bottle".to_string(),
        2 => "two green bottles".to_string(),
        3 => "three green bottles".to_string(),
        4 => "four green bottles".to_string(),
        5 => "five green bottles".to_string(),
        6 => "six green bottles".to_string(),
        7 => "seven green bottles".to_string(),
        8 => "eight green bottles".to_string(),
        9 => "nine green bottles".to_string(),
        10 => "ten green bottles".to_string(),
        _ => unreachable!(),
    }
}

fn bottles_cap(n: u32) -> String {
    match n {
        0 => "No green bottles".to_string(),
        1 => "One green bottle".to_string(),
        2 => "Two green bottles".to_string(),
        3 => "Three green bottles".to_string(),
        4 => "Four green bottles".to_string(),
        5 => "Five green bottles".to_string(),
        6 => "Six green bottles".to_string(),
        7 => "Seven green bottles".to_string(),
        8 => "Eight green bottles".to_string(),
        9 => "Nine green bottles".to_string(),
        10 => "Ten green bottles".to_string(),
        _ => unreachable!(),
    }
}

pub fn recite(start_bottles: u32, take_down: u32) -> String {
    let mut song = String::new();

    for bottles_left in (start_bottles - take_down + 1..=start_bottles).rev() {
        song.push_str(&format!(
            "{0} hanging on the wall,\n\
             {0} hanging on the wall,\n\
             And if one green bottle should accidentally fall,\n\
             There'll be {1} hanging on the wall.",
            bottles_cap(bottles_left),
            bottles(bottles_left - 1)
        ));

        if bottles_left != start_bottles - take_down + 1 {
            song.push_str("\n\n");
        }
    }

    song
}
