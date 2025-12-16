use anyhow::*;
use aoc_2025::*;
use code_timing_macros::time_snippet;
use const_format::concatcp;
use std::fs::File;
use std::io::{BufRead, BufReader};

const DAY: &str = "05";
const INPUT_FILE: &str = concatcp!("input/", DAY, ".txt");

const TEST: &str = "\
3-5
10-14
16-20
12-18

1
5
8
11
17
32
";

fn main() -> Result<()> {
    start_day(DAY);

    //region Part 1
    println!("=== Part 1 ===");

    fn part1<R: BufRead>(reader: R) -> Result<usize> {
        let mut fresh_list = Vec::new();
        let mut answer = 0;
        for line in reader.lines().flatten().filter(|line| !line.is_empty()) {
            if line.contains('-') {
                let (start, end) = line.split_once('-').unwrap();
                let start = start.parse::<usize>().unwrap();
                let end = end.parse::<usize>().unwrap();
                fresh_list.push((start, end));
            } else {
                let ingredient = line.parse::<usize>().unwrap();
                let mut is_fresh = false;
                for (start, end) in &fresh_list {
                    if ingredient >= *start && ingredient <= *end {
                        is_fresh = true;
                        break;
                    }
                }
                if is_fresh {
                    answer += 1;
                }
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

    fn part2<R: BufRead>(reader: R) -> Result<usize> {
        let mut fresh_list = Vec::new();
        for line in reader.lines().flatten().filter(|line| !line.is_empty()) {
            if line.contains('-') {
                let (start, end) = line.split_once('-').unwrap();
                let start = start.parse::<usize>().unwrap();
                let end = end.parse::<usize>().unwrap();
                fresh_list.push((start, end));
            } else {
                break;
            }
        }
        // sort intervals by start value
        fresh_list.sort_by_key(|k| k.0);
        // merge intervals
        let mut merged_list = Vec::new();
        let mut current_start = fresh_list[0].0;
        let mut current_end = fresh_list[0].1;
        for i in 1..fresh_list.len() {
            let (start, end) = fresh_list[i];
            if start <= current_end {
                // overlapping intervals, extend the end if needed
                if end > current_end {
                    current_end = end;
                }
            } else {
                // non-overlapping interval, push the current interval and reset
                merged_list.push((current_start, current_end));
                current_start = start;
                current_end = end;
            }
        }
        // push the last interval
        merged_list.push((current_start, current_end));
        fresh_list = merged_list;
        Ok(fresh_list.iter().map(|range| range.1 - range.0 + 1).sum())
    }

    assert_eq!(14, part2(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part2(input_file)?);
    println!("Result = {}", result);
    //endregion

    Ok(())
}
