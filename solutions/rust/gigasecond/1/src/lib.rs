use time::{PrimitiveDateTime as DateTime, SignedDuration};

// Returns a DateTime one billion seconds after start.
pub fn after(start: DateTime) -> DateTime {
    let t_million: SignedDuration = 1_000_000_000 * SignedDuration::SECOND;
    start.checked_add(t_million).unwrap()
}
