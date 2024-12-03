use std::{
    fs::File,
    io::{BufRead, BufReader},
    time::Instant,
};

pub fn run() {
    println!("*** Advent of Code 2024 ***");
    println!("--- Day 3: Mull It Over ---");
    // Quite wordy solutions, but I bet that they are still better than regex.
    part1();
    part2();
}

// Reading character by character:
// https://stackoverflow.com/questions/35385703/read-file-character-by-character-in-rust/37189758
fn part1() {
    let mut f =
        BufReader::new(File::open("input/day-03.txt").expect("failed to read the input file"));

    let start = Instant::now();

    let mut buf = Vec::<u8>::new();
    let mut instruction = String::new();
    let mut reading_first = false;
    let mut first_s = String::new();
    let mut reading_second = false;
    let mut second_s = String::new();
    let mut sum = 0;
    while f.read_until(b'\n', &mut buf).expect("read_until failed") != 0 {
        let s = String::from_utf8(buf).expect("from_utf8 failed");
        for c in s.chars() {
            if (instruction == "" && c == 'm')
                || (instruction == "m" && c == 'u')
                || (instruction == "mu" && c == 'l')
            {
                instruction.push(c);
                continue;
            }

            if instruction == "mul" && c == '(' {
                instruction.push(c);
                reading_first = true;
                continue;
            }

            if instruction == "mul(" && reading_first && c.is_digit(10) {
                first_s.push(c);
                continue;
            }

            if instruction == "mul(" && reading_first && !first_s.is_empty() && c == ',' {
                instruction.push(c);
                reading_first = false;
                reading_second = true;
                continue;
            }

            if instruction == "mul(," && reading_second && c.is_digit(10) {
                second_s.push(c);
                continue;
            }

            if instruction == "mul(," && reading_second && !second_s.is_empty() && c == ')' {
                instruction.push(c);
                let first: i32 = first_s.parse().unwrap();
                let second: i32 = second_s.parse().unwrap();
                sum += first * second;
            }

            // If we get here, the sequence is invalid and we should reset
            // everything. By clearing the strings, they keep their capacity,
            // which I think is good.
            instruction.clear();
            reading_first = false;
            first_s.clear();
            reading_second = false;
            second_s.clear();
        }
        buf = s.into_bytes();
        buf.clear();
    }

    let d = start.elapsed();

    println!("Part 1: the sum of the multiplications is {}", sum);
    println!("Part 1 ran for {:?}", d);
}

fn part2() {
    let mut f =
        BufReader::new(File::open("input/day-03.txt").expect("failed to read the input file"));

    let start = Instant::now();

    let mut buf = Vec::<u8>::new();
    let mut enabled = true;
    let mut instruction = String::new();
    let mut reading_first = false;
    let mut first_s = String::new();
    let mut reading_second = false;
    let mut second_s = String::new();
    let mut sum = 0;
    while f.read_until(b'\n', &mut buf).expect("read_until failed") != 0 {
        let s = String::from_utf8(buf).expect("from_utf8 failed");
        for c in s.chars() {
            if enabled {
                if c == 'd' || c == 'm' {
                    instruction.clear();
                    instruction.push(c);
                    reading_first = false;
                    first_s.clear();
                    reading_second = false;
                    second_s.clear();
                    continue;
                }

                if (instruction == "d" && c == 'o')
                    || (instruction == "do" && c == 'n')
                    || (instruction == "don" && c == '\'')
                    || (instruction == "don'" && c == 't')
                    || (instruction == "don't" && c == '(')
                {
                    instruction.push(c);
                    continue;
                }

                if instruction == "don't(" && c == ')' {
                    enabled = false;
                    instruction.clear();
                    reading_first = false;
                    first_s.clear();
                    reading_second = false;
                    second_s.clear();
                    continue;
                }

                if (instruction == "m" && c == 'u') || (instruction == "mu" && c == 'l') {
                    instruction.push(c);
                    continue;
                }

                if instruction == "mul" && c == '(' {
                    instruction.push(c);
                    reading_first = true;
                    continue;
                }

                if instruction == "mul(" && reading_first && c.is_digit(10) {
                    first_s.push(c);
                    continue;
                }

                if instruction == "mul(" && reading_first && !first_s.is_empty() && c == ',' {
                    instruction.push(c);
                    reading_first = false;
                    reading_second = true;
                    continue;
                }

                if instruction == "mul(," && reading_second && c.is_digit(10) {
                    second_s.push(c);
                    continue;
                }

                if instruction == "mul(," && reading_second && !second_s.is_empty() && c == ')' {
                    instruction.push(c);
                    let first: i32 = first_s.parse().unwrap();
                    let second: i32 = second_s.parse().unwrap();
                    sum += first * second;
                }

                // If we get here, the sequence is invalid and we should reset
                // everything. By clearing the strings, they keep their capacity,
                // which I think is good.
                instruction.clear();
                reading_first = false;
                first_s.clear();
                reading_second = false;
                second_s.clear();
            } else {
                if (instruction == "" && c == 'd')
                    || (instruction == "d" && c == 'o')
                    || (instruction == "do" && c == '(')
                {
                    instruction.push(c);
                    continue;
                }

                if instruction == "do(" && c == ')' {
                    enabled = true;
                }

                instruction.clear();
                reading_first = false;
                first_s.clear();
                reading_second = false;
                second_s.clear();
            }
        }
        buf = s.into_bytes();
        buf.clear();
    }

    let d = start.elapsed();

    println!("Part 2: the sum of the multiplications is {}", sum);
    println!("Part 2 ran for {:?}", d);
}
