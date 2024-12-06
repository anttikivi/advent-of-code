use std::env;

mod day_01;
mod day_02;
mod day_03;
mod day_04;

fn main() {
    let args: Vec<String> = env::args().collect();
    assert_eq!(
        args.len(),
        2,
        "You must pass only the day number as an argument! Now got the following arguments: {args:#?}"
    );
    let day_to_run = args[1].parse::<i32>().unwrap();
    match day_to_run {
        1 => day_01::run(),
        2 => day_02::run(),
        3 => day_03::run(),
        4 => day_04::run(),
        _ => println!("No solution for day {day_to_run}"),
    }
}
