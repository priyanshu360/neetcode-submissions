use std::collections::HashSet;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &'a [&str]) -> HashSet<&'a str> {
    possible_anagrams
        .iter()
        .copied()
        .filter(|candidate| {
            let mut w: Vec<char> = word.to_lowercase().chars().collect();
            let mut c: Vec<char> = candidate.to_lowercase().chars().collect();
            if w == c {
                return false;
            };
            w.sort();
            c.sort();
            return w == c;
        })
        .collect()
}
