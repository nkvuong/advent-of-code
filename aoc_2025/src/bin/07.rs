use anyhow::*;
use aoc_2025::*;
use code_timing_macros::time_snippet;
use const_format::concatcp;
use std::collections::HashSet;
use std::fs::File;
use std::io::{BufRead, BufReader};

const DAY: &str = "07";
const INPUT_FILE: &str = concatcp!("input/", DAY, ".txt");

const TEST: &str = "\
.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
";

fn main() -> Result<()> {
    start_day(DAY);

    //region Part 1
    println!("=== Part 1 ===");

    fn part1<R: BufRead>(reader: R) -> Result<usize> {
        let mut beams: HashSet<usize> = HashSet::new();
        let mut answer = 0;
        for line in reader.lines().flatten() {
            let mut new_beams: HashSet<usize> = HashSet::new();
            if beams.is_empty() {
                new_beams.insert(line.find('S').unwrap());
            } else {
                if !line.contains('^') {
                    continue; // skip rows without splits
                }
                for beam in beams {
                    let ch = line.chars().nth(beam).unwrap();
                    match ch {
                        '.' => {new_beams.insert(beam);}, // keep the beam
                        '^' => {
                            // split the beam
                            if beam > 0 {
                                new_beams.insert(beam - 1);
                            }
                            if beam + 1 < line.len() {
                                new_beams.insert(beam + 1);
                            }
                            if beam > 0 && beam + 1 < line.len() {
                                answer += 1; // count splits
                            }
                        }
                        _ => {} // beam is lost
                    }
                }
            }
            beams = new_beams;
        }
        Ok(answer)
    }

    assert_eq!(21, part1(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part1(input_file)?);
    println!("Result = {}", result);
    //endregion

    //region Part 2
    println!("\n=== Part 2 ===");

    fn part2<R: BufRead>(reader: R) -> Result<usize> {
        let lines = reader.lines().flatten().collect::<Vec<String>>();
        let first_line = &lines[0];
        let start_index = first_line.find('S').unwrap();
        let width = first_line.len();
        let mut timelines = vec![1; width];
        // process from bottom to top
        // each '^' contains the sum of timelines from left and right children
        for line in lines.iter().rev() {
            if !line.contains('^') {
                continue; // skip rows without splits
            }
            for (i, ch) in line.chars().enumerate() {
                if ch == '^' {
                    timelines[i] = timelines[i - 1] + timelines[i + 1];
                }
            }
        }
        Ok(timelines[start_index])
    }

    assert_eq!(40, part2(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part2(input_file)?);
    println!("Result = {}", result);
    //endregion

    Ok(())
}
