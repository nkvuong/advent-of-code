use anyhow::*;
use aoc_2025::*;
use code_timing_macros::time_snippet;
use const_format::concatcp;
use std::collections::{HashMap, HashSet};
use std::fs::File;
use std::io::{BufRead, BufReader};

const DAY: &str = "08";
const INPUT_FILE: &str = concatcp!("input/", DAY, ".txt");

const TEST: &str = "\
162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689
";

struct Box {
    x: isize,
    y: isize,
    z: isize,
}

impl Box {
    fn from_str(s: String) -> Self {
        let parts: Vec<isize> = s
            .split(',')
            .map(|num_str| num_str.parse::<isize>().unwrap())
            .collect();
        Box {
            x: parts[0],
            y: parts[1],
            z: parts[2],
        }
    }
}

fn distance(a: &Box, b: &Box) -> f32 {
    let dx = a.x - b.x;
    let dy: isize = a.y - b.y;
    let dz: isize = a.z - b.z;
    ((dx * dx + dy * dy + dz * dz) as f32).sqrt()
}

fn get_distances(boxes: &Vec<Box>) -> Vec<(usize, usize, f32)> {
    let num_boxes = boxes.len();
    let mut distances: Vec<(usize, usize, f32)> =
        Vec::with_capacity(num_boxes * (num_boxes - 1) / 2);
    for i in 0..num_boxes {
        for j in (i + 1)..num_boxes {
            let dist = distance(&boxes[i], &boxes[j]);
            distances.push((i, j, dist));
        }
    }
    distances.sort_by(|a, b| a.2.partial_cmp(&b.2).unwrap());
    distances
}

fn main() -> Result<()> {
    start_day(DAY);

    //region Part 1
    println!("=== Part 1 ===");

    fn part1<R: BufRead>(reader: R, num_connections: usize) -> Result<usize> {
        let boxes: Vec<Box> = reader.lines().flatten().map(Box::from_str).collect();
        let num_boxes = boxes.len();
        let distances = get_distances(&boxes);
        // make the closest connections
        let mut connections: HashMap<usize, HashSet<usize>> = HashMap::with_capacity(num_boxes);
        for k in 0..num_connections {
            let (first_box, second_box, _) = distances[k];
            connections.entry(first_box).or_default().insert(second_box);
            connections.entry(second_box).or_default().insert(first_box);
        }
        // get all connected circuits and the size of each circuit
        let mut visited: HashSet<usize> = HashSet::with_capacity(num_boxes);
        let mut circuit_sizes: Vec<usize> = Vec::with_capacity(num_boxes);
        for start_box in 0..num_boxes {
            if visited.contains(&start_box) {
                continue;
            }
            let mut size = 0;
            let mut stack: Vec<usize> = Vec::with_capacity(num_boxes);
            stack.push(start_box);
            while let Some(current_box) = stack.pop() {
                if visited.contains(&current_box) {
                    continue;
                }
                visited.insert(current_box);
                size += 1;
                if let Some(neighbors) = connections.get(&current_box) {
                    for &neighbor in neighbors {
                        if !visited.contains(&neighbor) {
                            stack.push(neighbor);
                        }
                    }
                }
            }
            circuit_sizes.push(size);
        }
        // sort circuit sizes descending
        circuit_sizes.sort_by(|a, b| b.cmp(a));
        Ok(circuit_sizes[0] * circuit_sizes[1] * circuit_sizes[2])
    }

    assert_eq!(40, part1(BufReader::new(TEST.as_bytes()), 10)?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part1(input_file, 1000)?);
    println!("Result = {}", result);
    //endregion

    //region Part 2
    println!("\n=== Part 2 ===");

    fn part2<R: BufRead>(reader: R) -> Result<isize> {
        let boxes: Vec<Box> = reader.lines().flatten().map(Box::from_str).collect();
        let num_boxes = boxes.len();
        let distances = get_distances(&boxes);
        let mut connections: HashMap<usize, HashSet<usize>> = HashMap::with_capacity(num_boxes);
        let mut connected_set: HashSet<usize> = HashSet::with_capacity(num_boxes);
        connected_set.insert(0); // start from box 0

        // keep adding connections until all boxes are connected
        for (first_box, second_box, _) in distances {
            connections.entry(first_box).or_default().insert(second_box);
            connections.entry(second_box).or_default().insert(first_box);
            // if either box is connected, add all its connections to connected_set
            if connected_set.contains(&first_box) || connected_set.contains(&second_box) {
                // do a DFS to find all connected boxes
                let mut stack: Vec<usize> = vec![first_box, second_box];
                while let Some(current_box) = stack.pop() {
                    if connected_set.contains(&current_box) {
                        continue;
                    }
                    connected_set.insert(current_box);
                    if let Some(neighbors) = connections.get(&current_box) {
                        for &neighbor in neighbors {
                            if !connected_set.contains(&neighbor) {
                                stack.push(neighbor);
                            }
                        }
                    }
                }
            }
            // check if we can reach all boxes from box 0
            if connected_set.len() == num_boxes {
                return Ok(boxes[first_box].x * boxes[second_box].x);
            }
        }
        Err(anyhow!("Could not connect all boxes"))
    }

    assert_eq!(25272, part2(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part2(input_file)?);
    println!("Result = {}", result);
    //endregion

    Ok(())
}
