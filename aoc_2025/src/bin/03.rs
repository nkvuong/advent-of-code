use anyhow::*;
use aoc_2025::*;
use code_timing_macros::time_snippet;
use const_format::concatcp;
use std::fs::File;
use std::io::{BufRead, BufReader};

const DAY: &str = "03";
const INPUT_FILE: &str = concatcp!("input/", DAY, ".txt");

const TEST: &str = "\
987654321111111
811111111111119
234234234234278
818181911112111
";

fn max_joltage(line: String, num_chars: usize) -> usize {
    // input is decimal digits
    let bytes = line.as_bytes();
    let mut start = 0;
    let mut answer = 0usize;

    // Greedily pick the maximum digit from the remaining window,
    // ensuring enough characters are left to reach `num_chars`.
    // this works because the max digit at each step is always optimal.
    for remaining in (0..num_chars).rev() {
        // leave enough characters to reach num_chars
        let end = bytes.len() - remaining;

        let mut max_digit = b'0';
        // start after the last picked digit
        let mut max_pos = start;
        for i in start..end {
            let d = bytes[i];
            if d > max_digit {
                max_digit = d;
                max_pos = i;
            }
        }

        answer = answer * 10 + (max_digit - b'0') as usize;
        start = max_pos + 1;
    }

    answer
}

fn main() -> Result<()> {
    start_day(DAY);

    //region Part 1
    println!("=== Part 1 ===");

    fn part1<R: BufRead>(reader: R) -> Result<usize> {
        let answer = reader
            .lines()
            .flatten()
            .map(|line| max_joltage(line, 2))
            .sum();
        Ok(answer)
    }

    assert_eq!(357, part1(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part1(input_file)?);
    println!("Result = {}", result);
    //endregion

    //region Part 2
    println!("\n=== Part 2 ===");

    fn part2<R: BufRead>(reader: R) -> Result<usize> {
        let answer = reader
            .lines()
            .flatten()
            .map(|line| max_joltage(line, 12))
            .sum();
        Ok(answer)
    }

    assert_eq!(3121910778619, part2(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part2(input_file)?);
    println!("Result = {}", result);
    //endregion

    Ok(())
}
