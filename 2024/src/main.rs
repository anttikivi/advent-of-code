use std::env;

mod day1;
mod day2;

fn main() {
    let args: Vec<String> = env::args().collect();
    assert_eq!(
        args.len(),
        2,
        "You must pass only the day number as an argument! Now got the following arguments: {args:#?}"
    );
    let day_to_run = args[1].parse::<i32>().unwrap();
    match day_to_run {
        1 => day1::run(),
        2 => day2::run(),
        _ => println!("No solution for day {day_to_run}"),
    }
}
