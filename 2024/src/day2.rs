use std::{
    fs::File,
    io::{BufRead, BufReader},
    time::Instant,
};

pub fn run() {
    println!("***    Advent of Code 2024   ***");
    println!("--- Day 2: Red-Nosed Reports ---");
    run_part1();
    run_part2();
}

fn is_report_safe(vec: &Vec<i32>) -> bool {
    // The parts must have a difference of at least one, so if the second
    // is not strictly greater than the first, the numbers must be
    // decresing.
    let increasing = vec[0] < vec[1];
    for i in 1..vec.len() {
        let diff = vec[i].abs_diff(vec[i - 1]);
        if diff < 1 || diff > 3 {
            return false;
        }
        if increasing && vec[i - 1] >= vec[i] {
            return false;
        }
        if !increasing && vec[i - 1] < vec[i] {
            return false;
        }
    }

    return true;
}

fn run_part1() {
    let lines =
        BufReader::new(File::open("input/day-02.txt").expect("failed to read the input file"))
            .lines()
            .filter_map(|r| r.ok());

    let start = Instant::now();

    let count = lines
        .map(|line| {
            line.split(" ")
                .map(|s| s.to_string().parse::<i32>().unwrap())
                .collect()
        })
        .filter(|l: &Vec<i32>| is_report_safe(&l))
        .count();

    let d = start.elapsed();

    println!("Part 1: there are {} safe reports in total", count);
    println!("Part 1 ran for {:?}", d);
}

fn run_part2() {
    let lines: Vec<String> =
        BufReader::new(File::open("input/day-02.txt").expect("failed to read the input file"))
            .lines()
            .filter_map(|result| result.ok())
            .collect();

    let start = Instant::now();

    let total = lines.len();
    let not_safe: Vec<Vec<i32>> = lines
        .into_iter()
        .map(|line| {
            line.split(" ")
                .map(|s| s.to_string().parse::<i32>().unwrap())
                .collect()
        })
        .filter(|l| !is_report_safe(l))
        .collect();
    // The count starts at the number of reports that are safe without
    // dampening. There is no need to handle those again. Additionally, now I
    // know that the rest of the reports need special handling in order to be
    // safe.
    let mut count = total - not_safe.len();
    for mut report in not_safe {
        for i in 0..report.len() {
            let num = report.remove(i);
            if is_report_safe(&report) {
                count += 1;
                break;
            }
            report.insert(i, num);
        }
    }

    let d = start.elapsed();

    println!("Part 2: there are {} safe reports in total", count);
    println!("Part 2 ran for {:?}", d);
}
