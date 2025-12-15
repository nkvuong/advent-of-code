use anyhow::*;
use aoc_2025::*;
use code_timing_macros::time_snippet;
use const_format::concatcp;
use std::fs::File;
use std::io::{BufRead, BufReader};

const DAY: &str = "01";
const INPUT_FILE: &str = concatcp!("input/", DAY, ".txt");

const TEST: &str = "\
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
";

fn main() -> Result<()> {
    start_day(DAY);

    //region Part 1
    println!("=== Part 1 ===");

    fn part1<R: BufRead>(reader: R) -> Result<usize> {
        let rotations: Vec<isize> = reader
            .lines()
            .flatten()
            .map(|line| {
                let (turn, value) = line.split_at(1);
                let value = match turn {
                    "L" => -value.parse::<isize>().unwrap(), // left turn
                    "R" => value.parse::<isize>().unwrap(),  // right turn
                    _ => 0,
                };
                value
            })
            .collect();
        let mut pos: isize = 50;
        let mut answer = 0;
        for rotation in rotations {
            pos = (pos + rotation + 100) % 100;
            if pos == 0 {
                answer += 1;
            }
        }
        Ok(answer)
    }

    assert_eq!(3, part1(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part1(input_file)?);
    println!("Result = {}", result);
    //endregion

    //region Part 2
    println!("\n=== Part 2 ===");

    fn part2<R: BufRead>(reader: R) -> Result<isize> {
        let rotations: Vec<isize> = reader
            .lines()
            .flatten()
            .map(|line| {
                let (turn, value) = line.split_at(1);
                let value = match turn {
                    "L" => -value.parse::<isize>().unwrap(), // left turn
                    "R" => value.parse::<isize>().unwrap(),  // right turn
                    _ => 0,
                };
                value
            })
            .collect();
        let mut pos: isize = 50;
        let mut answer = 0;
        for rotation in rotations {
            answer += rotation.abs() / 100; // number of full loops
            pos = pos + rotation % 100;
            // check if we crossed the 0 point on a right turn
            if pos >= 100 {
                answer += 1;
            }
            // check if we crossed the 0 point on a left turn. exclude when we start at 0
            if pos <= 0 && pos != rotation % 100 {
                answer += 1;
            }
            pos = (pos + 100) % 100;
        }
        Ok(answer)
    }

    assert_eq!(6, part2(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part2(input_file)?);
    println!("Result = {}", result);
    //endregion

    Ok(())
}
