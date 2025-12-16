use anyhow::*;
use aoc_2025::*;
use code_timing_macros::time_snippet;
use const_format::concatcp;
use std::fs::File;
use std::io::{BufRead, BufReader};

const DAY: &str = "06";
const INPUT_FILE: &str = concatcp!("input/", DAY, ".txt");

const TEST: &str = "\
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
";

fn main() -> Result<()> {
    start_day(DAY);

    //region Part 1
    println!("=== Part 1 ===");

    fn part1<R: BufRead>(reader: R) -> Result<usize> {
        let mut numbers: Vec<Vec<usize>> = Vec::new();
        let mut answer = 0;
        for line in reader.lines().flatten() {
            // filter just the operator
            if line.contains('*') {
                for (col, op) in line.chars().filter(|c| *c == '*' || *c == '+').enumerate() {
                    let mut col_result = if op == '*' { 1 } else { 0 };
                    for row in 0..numbers.len() {
                        let val = numbers[row][col];
                        if op == '*' {
                            col_result *= val;
                        } else {
                            col_result += val;
                        }
                    }
                    answer += col_result;
                }
            } else {
                // parse numbers
                numbers.push(
                    line.split_whitespace()
                        .map(|num_str| num_str.parse::<usize>().unwrap())
                        .collect(),
                );
            }
        }
        Ok(answer)
    }

    assert_eq!(4277556, part1(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part1(input_file)?);
    println!("Result = {}", result);
    //endregion

    //region Part 2
    println!("\n=== Part 2 ===");

    fn part2<R: BufRead>(reader: R) -> Result<usize> {
        let mut answer = 0;
        let input: Vec<Vec<char>> = reader
            .lines()
            .flatten()
            .map(|line| line.chars().rev().collect())
            .collect();
        let num_rows = input.len();
        let num_cols = input[0].len();
        let mut numbers: Vec<usize> = Vec::new();
        for col in 0..num_cols {
            let mut num = 0;
            for row in 0..num_rows {
                // at the last row, push the number
                if row == num_rows - 1 && num > 0 {
                    numbers.push(num);
                }
                let ch = input[row][col];
                if ch.is_numeric() {
                    let digit = ch.to_digit(10).unwrap() as usize;
                    num = num * 10 + digit;
                } else if ch == '*' {
                    answer += numbers.iter().product::<usize>();
                    numbers.clear();
                } else if ch == '+' {
                    answer += numbers.iter().sum::<usize>();
                    numbers.clear();
                }
            }
        }
        Ok(answer)
    }

    assert_eq!(3263827, part2(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part2(input_file)?);
    println!("Result = {}", result);
    //endregion

    Ok(())
}
