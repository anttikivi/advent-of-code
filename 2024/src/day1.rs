use std::{
    fs::File,
    io::{BufRead, BufReader},
};

const SEPARATOR: &str = "   ";

pub fn run() {
    println!("***    Advent of Code 2024    ***");
    println!("--- Day 1: Historian Hysteria ---");

    let input_file = File::open("input/day-01.txt").expect("failed to read the input file");
    let reader = BufReader::new(input_file);

    let mut left_nums: Vec<i32> = Vec::with_capacity(1000);
    let mut right_nums: Vec<i32> = Vec::with_capacity(1000);

    for line in reader.lines().filter_map(|result| result.ok()) {
        let (left_s, right_s) = line.split_once(SEPARATOR).unwrap();
        let left = left_s.to_string().parse::<i32>().unwrap();
        let right = right_s.to_string().parse::<i32>().unwrap();
        left_nums.push(left);
        right_nums.push(right);
    }

    left_nums.sort_unstable();
    right_nums.sort_unstable();

    let mut sum = 0;
    for i in 0..left_nums.len() {
        let diff = left_nums[i].abs_diff(right_nums[i]);
        sum += diff
    }

    println!("Part 1: the total distance between the two lists is {sum}");
}
