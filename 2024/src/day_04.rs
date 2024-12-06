use std::{
    collections::HashMap,
    fs::File,
    io::{BufRead, BufReader},
    time::Instant,
};

pub fn run() {
    println!("*** Advent of Code 2024 ***");
    println!("--- Day 4: Ceres Search ---");
    part1();
    part2();
}

fn part1() {
    let mut f =
        BufReader::new(File::open("input/day-04.txt").expect("failed to read the input file"));

    let start = Instant::now();

    let mut buf = Vec::<u8>::new();
    let mut x = 0;
    let mut y = 0;
    let mut max_x = 0;
    let mut max_y = 0;
    let mut chars: HashMap<(i32, i32), char> = HashMap::new();
    while f.read_until(b'\n', &mut buf).expect("read_until failed") != 0 {
        let s = String::from_utf8(buf).expect("from_utf8 failed");
        for c in s.chars() {
            // println!("Insert {} at {}, {}", c, x, y);
            chars.insert((x, y), c);
            x += 1;
        }

        if x > max_x {
            max_x = x;
        }

        x = 0;
        y += 1;

        if y > max_y {
            max_y = y;
        }

        buf = s.into_bytes();
        buf.clear();
    }

    let mut checked: Vec<(i32, i32)> = Vec::new();
    let mut sum = 0;
    for y in 0..max_y {
        for x in 0..max_x {
            let coord = (x, y);
            if *chars.get(&coord).unwrap() == 'X' && !checked.contains(&coord) {
                sum += count_around(x, y, max_x, max_y, &chars);
                checked.push(coord);
            }
        }
    }

    let d = start.elapsed();

    println!("Part 1: XMAS appears {sum} times");
    println!("Part 1 ran for {:?}", d);
}

fn count_around(x: i32, y: i32, max_x: i32, max_y: i32, chars: &HashMap<(i32, i32), char>) -> i32 {
    let x_diff = [
        [0, 0, 0],
        [1, 1, 1],
        [1, 1, 1],
        [1, 1, 1],
        [0, 0, 0],
        [-1, -1, -1],
        [-1, -1, -1],
        [-1, -1, -1],
    ];
    let y_diff = [
        [-1, -1, -1],
        [-1, -1, -1],
        [0, 0, 0],
        [1, 1, 1],
        [1, 1, 1],
        [1, 1, 1],
        [0, 0, 0],
        [-1, -1, -1],
    ];

    let mut sum = 0;

    for i in 0..8 {
        let mut coord = (x, y);
        let mut s = String::new();
        s.push(*chars.get(&coord).unwrap());
        // print!("{}", *chars.get(&(x, y)).unwrap());
        for j in 0..3 {
            coord.0 += x_diff[i][j];
            coord.1 += y_diff[i][j];
            if coord.0 < 0 || coord.1 < 0 || coord.0 >= max_x || coord.1 >= max_y {
                break;
            }
            s.push(*chars.get(&coord).unwrap());
        }
        if s == String::from("XMAS") {
            sum += 1;
        }
    }
    return sum;
}

fn part2() {
    let mut f =
        BufReader::new(File::open("input/day-04.txt").expect("failed to read the input file"));

    let start = Instant::now();

    let mut buf = Vec::<u8>::new();
    let mut x = 0;
    let mut y = 0;
    let mut max_x = 0;
    let mut max_y = 0;
    let mut chars: HashMap<(i32, i32), char> = HashMap::new();
    while f.read_until(b'\n', &mut buf).expect("read_until failed") != 0 {
        let s = String::from_utf8(buf).expect("from_utf8 failed");
        for c in s.chars() {
            chars.insert((x, y), c);
            x += 1;
        }

        if x > max_x {
            max_x = x;
        }

        x = 0;
        y += 1;

        if y > max_y {
            max_y = y;
        }

        buf = s.into_bytes();
        buf.clear();
    }

    let mut checked: Vec<(i32, i32)> = Vec::new();
    let mut sum = 0;
    for y in 0..max_y {
        for x in 0..max_x {
            let coord = (x, y);
            if *chars.get(&coord).unwrap() == 'A' && !checked.contains(&coord) {
                if has_x_mas(x, y, max_x, max_y, &chars) {
                    sum += 1;
                }
                checked.push(coord);
            }
        }
    }

    let d = start.elapsed();

    println!("Part 2: X-MAS appears {sum} times");
    println!("Part 2 ran for {:?}", d);
}

fn has_x_mas(x: i32, y: i32, max_x: i32, max_y: i32, chars: &HashMap<(i32, i32), char>) -> bool {
    let x_diff = [[-1, 1], [-1, 1]];
    let y_diff = [[-1, 1], [1, -1]];

    let diffs = (x_diff[0], y_diff[0]);
    let mut s = String::new();
    for i in 0..2 {
        let current_x = x + diffs.0[i];
        let current_y = y + diffs.1[i];
        if current_x < 0 || current_y < 0 || current_x >= max_x || current_y >= max_y {
            return false;
        }
        s.push(*chars.get(&(current_x, current_y)).unwrap());
    }

    if s != String::from("MS") && s != String::from("SM") {
        return false;
    }

    let diffs = (x_diff[1], y_diff[1]);
    let mut s = String::new();
    for i in 0..2 {
        let current_x = x + diffs.0[i];
        let current_y = y + diffs.1[i];
        if current_x < 0 || current_y < 0 || current_x >= max_x || current_y >= max_y {
            return false;
        }
        s.push(*chars.get(&(current_x, current_y)).unwrap());
    }

    if s != String::from("MS") && s != String::from("SM") {
        return false;
    }

    return true;
}
