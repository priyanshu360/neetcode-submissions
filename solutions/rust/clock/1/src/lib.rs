use core::fmt;

#[derive(PartialEq, Debug, Clone, Eq)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let total = (hours * 60 + minutes).rem_euclid(24 * 60);

        Self {
            hours: total / 60,
            minutes: total % 60,
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Self::new(self.hours, self.minutes + minutes)
    }
}

// impl PartialEq for Clock {
//     fn eq(&self, other: &Self) -> bool {
//         self.hours == other.hours && self.minutes == other.minutes
//     }
// }

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.hours, self.minutes)
    }
}
