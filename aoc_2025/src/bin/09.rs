use anyhow::*;
use aoc_2025::*;
use code_timing_macros::time_snippet;
use const_format::concatcp;
use std::fs::File;
use std::io::{BufRead, BufReader};

const DAY: &str = "09";
const INPUT_FILE: &str = concatcp!("input/", DAY, ".txt");

const TEST: &str = "\
7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
";

struct Point {
    x: isize,
    y: isize,
}

impl Point {
    fn from_str(s: String) -> Self {
        let parts: Vec<isize> = s
            .split(',')
            .map(|num_str| num_str.parse::<isize>().unwrap())
            .collect();
        Point {
            x: parts[0],
            y: parts[1],
        }
    }
    fn area(&self, other: &Point) -> usize {
        let width = (self.x - other.x).abs() + 1;
        let height = (self.y - other.y).abs() + 1;
        (width * height) as usize
    }
}

fn rectangle_in_polygon(point: &Point, other: &Point, polygon: &Vec<Point>) -> bool {
    // since the polygon is axis-aligned, we can just check that the polygon edges
    // do not intersect the rectangle, by comparing the min/max x/y of the rectangle to each edge
    let xmax = other.x.max(point.x);
    let ymax = other.y.max(point.y);
    let xmin = other.x.min(point.x);
    let ymin = other.y.min(point.y);
    polygon
        .iter()
        // get all edges by zipping each point with the next, wrapping around
        .zip(polygon.iter().cycle().skip(1).take(polygon.len()))
        .all(|(p1, p2)| {
            (p1.y >= ymax && p2.y >= ymax)
                || (p1.y <= ymin && p2.y <= ymin)
                || (p1.x <= xmin && p2.x <= xmin)
                || (p1.x >= xmax && p2.x >= xmax)
        })
}

fn main() -> Result<()> {
    start_day(DAY);

    //region Part 1
    println!("=== Part 1 ===");

    fn part1<R: BufRead>(reader: R) -> Result<usize> {
        let points: Vec<Point> = reader.lines().flatten().map(Point::from_str).collect();
        let mut max_area = 0;
        for i in 0..points.len() {
            for j in (i + 2)..points.len() {
                // since edges are axis-aligned, skip adjacent points
                max_area = max_area.max(points[i].area(&points[j]));
            }
        }
        Ok(max_area)
    }

    assert_eq!(50, part1(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part1(input_file)?);
    println!("Result = {}", result);
    //endregion

    //region Part 2
    println!("\n=== Part 2 ===");

    fn part2<R: BufRead>(reader: R) -> Result<usize> {
        let polygon: Vec<Point> = reader.lines().flatten().map(Point::from_str).collect();
        let mut max_area = 0;

        // Create list of (i, j, potential_area) and sort by potential_area descending
        let mut pairs: Vec<(usize, usize, usize)> = Vec::new();
        for i in 0..polygon.len() {
            for j in (i + 2)..polygon.len() {
                pairs.push((i, j, polygon[i].area(&polygon[j])));
            }
        }
        pairs.sort_by(|a, b| b.2.cmp(&a.2));

        // Check pairs in order of largest potential area first
        for (i, j, area) in pairs {
            // Early termination: if the rectangle is valid, it is the largest possible area
            if rectangle_in_polygon(&polygon[i], &polygon[j], &polygon) {
                max_area = area;
                break;
            }
        }
        Ok(max_area)
    }

    assert_eq!(24, part2(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part2(input_file)?);
    println!("Result = {}", result);
    //endregion

    Ok(())
}
