use anyhow::*;
use aoc_2025::*;
use code_timing_macros::time_snippet;
use const_format::concatcp;
use itertools::Itertools;
use std::fs::File;
use std::io::{BufRead, BufReader};

const DAY: &str = "02";
const INPUT_FILE: &str = concatcp!("input/", DAY, ".txt");

const TEST: &str = "\
11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124
";

fn main() -> Result<()> {
    start_day(DAY);

    //region Part 1
    println!("=== Part 1 ===");

    fn valid_id(id: usize, num_substr: usize) -> bool {
        let id_string = id.to_string();
        // string must be evenly divisible by num_substr
        if id_string.len() % num_substr != 0 {
            return true;
        }
        let substr_len = id_string.len() / num_substr;
        // check each substring matches the next, if there is a mismatch, then return
        for start in 0..num_substr - 1 {
            let substr = &id_string[start * substr_len..(start + 1) * substr_len];
            let next_substr = &id_string[(start + 1) * substr_len..(start + 2) * substr_len];
            if substr != next_substr {
                return true;
            }
        }
        // all substrings match, it is an invalid id
        false
    }

    fn part1<R: BufRead>(reader: R) -> Result<usize> {
        let ids: Vec<(usize, usize)> = reader
            .lines()
            .flatten()
            .map(|line| {
                let pairs = line.split(',');
                let ids = pairs.map(|pair| {
                    if !pair.contains('-') {
                        return (0, 0);
                    }
                    let (start, end) = pair.split_once('-').unwrap();
                    let start = start.parse::<usize>().unwrap();
                    let end = end.parse::<usize>().unwrap();
                    (start, end)
                });
                ids.collect_vec()
            })
            .concat();
        let mut answer = 0;
        for (start, end) in ids {
            for id in start..=end {
                if !valid_id(id, 2) {
                    answer += id;
                }
            }
        }
        Ok(answer)
    }

    assert_eq!(1227775554, part1(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part1(input_file)?);
    println!("Result = {}", result);
    //endregion

    //region Part 2
    println!("\n=== Part 2 ===");

    fn part2<R: BufRead>(reader: R) -> Result<usize> {
        let ids: Vec<(usize, usize)> = reader
            .lines()
            .flatten()
            .map(|line| {
                let pairs = line.split(',');
                let ids = pairs.map(|pair| {
                    if !pair.contains('-') {
                        return (0, 0);
                    }
                    let (start, end) = pair.split_once('-').unwrap();
                    let start = start.parse::<usize>().unwrap();
                    let end = end.parse::<usize>().unwrap();
                    (start, end)
                });
                ids.collect_vec()
            })
            .concat();
        let mut answer = 0;
        for (start, end) in ids {
            for id in start..=end {
                // check for all possible substring lengths
                for num_substr in 2..=id.to_string().len() {
                    if !valid_id(id, num_substr) {
                        answer += id;
                        break;
                    }
                }
            }
        }
        Ok(answer)
    }

    assert_eq!(4174379265, part2(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part2(input_file)?);
    println!("Result = {}", result);
    //endregion

    Ok(())
}
