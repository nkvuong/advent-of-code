use anyhow::*;
use aoc_2025::*;
use code_timing_macros::time_snippet;
use const_format::concatcp;
use itertools::Itertools;
use std::fs::File;
use std::io::{BufRead, BufReader};
use std::ops::{Add, Sub};
use std::usize::MAX;

const DAY: &str = "04";
const INPUT_FILE: &str = concatcp!("input/", DAY, ".txt");

const TEST: &str = "\
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
";

fn remove_papers(papers: &mut Vec<Vec<u8>>, num_times: usize) -> usize {
    if papers.is_empty() || num_times == 0 {
        return 0;
    }

    let h = papers.len();
    let w = papers[0].len();
    let dirs: [(isize, isize); 8] = [
        (-1, -1),
        (-1, 0),
        (-1, 1),
        (0, -1),
        (0, 1),
        (1, -1),
        (1, 0),
        (1, 1),
    ];

    // Precompute neighbor counts for all cells.
    let mut counts = vec![vec![0u8; w]; h];
    for r in 0..h {
        for c in 0..w {
            if papers[r][c] != b'@' {
                continue;
            }
            for (dr, dc) in dirs.iter() {
                let nr = r as isize + dr;
                let nc = c as isize + dc;
                if nr >= 0 && nr < h as isize && nc >= 0 && nc < w as isize {
                    let (nr, nc) = (nr as usize, nc as usize);
                    counts[nr][nc] = counts[nr][nc].add(1);
                }
            }
        }
    }

    // Initial removed: cells that would be removed in the first round.
    let mut removed: Vec<(usize, usize)> = Vec::new();
    for r in 0..h {
        for c in 0..w {
            if papers[r][c] == b'@' && counts[r][c] < 4 {
                removed.push((r, c));
            }
        }
    }

    let mut removed_total = 0;
    let mut to_remove = vec![vec![false; w]; h];
    let mut in_next = vec![vec![false; w]; h];

    for _ in 0..num_times {
        if removed.is_empty() {
            break;
        }

        removed_total += removed.len();
        for &(r, c) in &removed {
            to_remove[r][c] = true;
        }

        let mut next_removed: Vec<(usize, usize)> = Vec::new();

        // Update neighbor counts for cells not removed this round and collect next removed.
        for &(r, c) in &removed {
            for (dr, dc) in dirs.iter() {
                let nr = r as isize + dr;
                let nc = c as isize + dc;
                if nr < 0 || nr >= h as isize || nc < 0 || nc >= w as isize {
                    continue;
                }
                let (nr, nc) = (nr as usize, nc as usize);
                if to_remove[nr][nc] || papers[nr][nc] != b'@' {
                    continue;
                }
                counts[nr][nc] = counts[nr][nc].sub(1);
                if counts[nr][nc] < 4 && !in_next[nr][nc] {
                    in_next[nr][nc] = true;
                    next_removed.push((nr, nc));
                }
            }
        }

        // Apply removals for this round.
        for (r, c) in removed {
            papers[r][c] = b'.';
            to_remove[r][c] = false;
        }

        // Prepare next round.
        for &(r, c) in &next_removed {
            in_next[r][c] = false;
        }
        removed = next_removed;
    }

    removed_total
}

fn main() -> Result<()> {
    start_day(DAY);

    //region Part 1
    println!("=== Part 1 ===");

    fn part1<R: BufRead>(reader: R) -> Result<usize> {
        let mut papers = reader
            .lines()
            .flatten()
            .map(|line| line.into_bytes())
            .collect_vec();
        Ok(remove_papers(&mut papers, 1))
    }

    assert_eq!(13, part1(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part1(input_file)?);
    println!("Result = {}", result);
    //endregion

    //region Part 2
    println!("\n=== Part 2 ===");

    fn part2<R: BufRead>(reader: R) -> Result<usize> {
        let mut papers = reader
            .lines()
            .flatten()
            .map(|line| line.into_bytes())
            .collect_vec();
        Ok(remove_papers(&mut papers, MAX))
    }

    assert_eq!(43, part2(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part2(input_file)?);
    println!("Result = {}", result);
    //endregion

    Ok(())
}
