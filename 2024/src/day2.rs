use std::{
    fs::File,
    io::{BufRead, BufReader},
};

pub fn run() {
    println!("***    Advent of Code 2024   ***");
    println!("--- Day 2: Red-Nosed Reports ---");

    let input_file = File::open("input/day-02.txt").expect("failed to read the input file");
    let reader = BufReader::new(input_file);

    let mut count = 0;
    for line in reader.lines().filter_map(|r| r.ok()) {
        // println!("{}", line);
        let parts: Vec<i32> = line
            .split(" ")
            .map(|s| s.to_string().parse::<i32>().unwrap())
            .collect();
        // The parts must have a difference of at least one, so if the second
        // is not strictly greater than the first, the numbers must be
        // decresing.
        let increasing = parts[0] < parts[1];
        let mut valid = true;
        for i in 1..parts.len() {
            let diff = parts[i].abs_diff(parts[i - 1]);
            if diff < 1 || diff > 3 {
                valid = false;
                break;
            }
            if increasing && parts[i - 1] >= parts[i] {
                valid = false;
                break;
            }
            if !increasing && parts[i - 1] < parts[i] {
                valid = false;
                break;
            }
        }
        if valid {
            count += 1;
        }
    }

    println!("Part 1: there are {} safe reports in total", count);
}
