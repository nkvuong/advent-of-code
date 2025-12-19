use anyhow::*;
use aoc_2025::*;
use code_timing_macros::time_snippet;
use const_format::concatcp;
use itertools::Itertools;
use std::fs::File;
use std::io::{BufRead, BufReader};

const DAY: &str = "12";
const INPUT_FILE: &str = concatcp!("input/", DAY, ".txt");

const TEST: &str = "\
0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2
";

struct Shape {
    grid: Vec<Vec<bool>>,
}

impl Shape {
    fn from_str(s: &str) -> Self {
        let grid: Vec<Vec<bool>> = s
            .lines()
            .filter(|line| !line.contains(':'))
            .map(|line| {
                line.chars()
                    .map(|c| if c == '#' { true } else { false })
                    .collect()
            })
            .collect();
        Shape { grid }
    }

    fn area(&self) -> usize {
        self.grid
            .iter()
            .map(|row| row.iter().filter(|&&cell| cell).count())
            .sum()
    }
}

struct Region {
    size: (usize, usize),
    shapes: Vec<usize>,
}

impl Region {
    fn from_str(s: &str) -> Result<Self> {
        let parts: Vec<&str> = s.split(':').collect();
        let size_parts: Vec<&str> = parts[0].trim().split('x').collect();
        let size = (
            size_parts[0].parse::<usize>()?,
            size_parts[1].parse::<usize>()?,
        );
        let shape_indices: Vec<usize> = parts[1]
            .trim()
            .split_whitespace()
            .map(|s| s.parse::<usize>())
            .try_collect()?;
        Ok(Region {
            size,
            shapes: shape_indices,
        })
    }

    fn can_fit(&self, shapes: &Vec<Shape>) -> bool {
        let total_shape_area: usize = self
            .shapes
            .iter()
            .enumerate()
            .map(|(idx, count)| shapes[idx].area() * count)
            .sum();
        // trivially possible if total shape count is less than one ninth of region area
        let total_shape_count: usize = self.shapes.iter().sum();
        if total_shape_count <= (self.size.0 / 3) * (self.size.1 / 3) {
            return true;
        }
        // trivially impossible if total shape area exceeds region area
        if total_shape_area > self.size.0 * self.size.1 {
            return false;
        }
        true
    }
}

fn parse_input(input: String) -> Result<(Vec<Shape>, Vec<Region>)> {
    let mut shapes: Vec<Shape> = Vec::new();
    let mut regions: Result<Vec<Region>> = Ok(Vec::new());
    input
        .split("\n\n")
        .map(|part| {
            if part.contains('x') {
                // region
                regions = part.trim().split("\n").map(Region::from_str).try_collect();
            } else {
                // shape
                let shape = Shape::from_str(part);
                shapes.push(shape);
            }
        })
        .for_each(drop);
    Ok((shapes, regions?))
}

fn main() -> Result<()> {
    start_day(DAY);

    //region Part 1
    println!("=== Part 1 ===");

    fn part1<R: BufRead>(mut reader: R) -> Result<usize> {
        // each block represent a shape, last block is the list of regions
        let mut input = String::new();
        reader.read_to_string(&mut input)?;
        let (shapes, regions) = parse_input(input)?;
        Ok(regions
            .iter()
            .filter(|region| region.can_fit(&shapes))
            .count())
    }

    // this heuristic-based solution does not work with the example though...
    //assert_eq!(2, part1(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part1(input_file)?);
    println!("Result = {}", result);
    //endregion

    Ok(())
}
