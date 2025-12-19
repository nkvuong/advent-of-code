use anyhow::*;
use aoc_2025::*;
use code_timing_macros::time_snippet;
use const_format::concatcp;
use std::collections::HashMap;
use std::fs::File;
use std::io::{BufRead, BufReader};

const DAY: &str = "11";
const INPUT_FILE: &str = concatcp!("input/", DAY, ".txt");

const TEST1: &str = "\
aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out
";

const TEST2: &str = "\
svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out
";

struct Graph {
    connections: HashMap<String, Vec<String>>,
}

impl Graph {
    fn from_str(s: String) -> Self {
        let mut connections: HashMap<String, Vec<String>> = HashMap::new();
        for line in s.lines() {
            let parts: Vec<&str> = line.split(':').collect();
            let node = parts[0].trim().to_string();
            let edges: Vec<String> = if parts.len() > 1 {
                parts[1]
                    .trim()
                    .split_whitespace()
                    .map(|e| e.to_string())
                    .collect()
            } else {
                vec![]
            };
            connections.insert(node, edges);
        }
        Graph { connections }
    }

    fn dfs_with_required(
        &self,
        current: &str,
        finish: &str,
        required: &Vec<&str>,
        mut visited_required: Vec<bool>,
        memo: &mut HashMap<(String, Vec<bool>), usize>,
    ) -> usize {
        // mark current as visited if it's in required
        if let Some(pos) = required.iter().position(|&node| node == current) {
            visited_required[pos] = true;
        }
        // if reached finish, check if all required have been visited
        if current == finish {
            return usize::from(visited_required.iter().all(|&seen| seen));
        }

        let state_key = (current.to_string(), visited_required.clone());
        if let Some(&cached) = memo.get(&state_key) {
            return cached;
        }

        let count = if let Some(edges) = self.connections.get(current) {
            edges
                .iter()
                .map(|edge| {
                    self.dfs_with_required(edge, finish, required, visited_required.clone(), memo)
                })
                .sum()
        } else {
            0
        };

        memo.insert(state_key, count);
        count
    }
}

fn main() -> Result<()> {
    start_day(DAY);

    //region Part 1
    println!("=== Part 1 ===");

    fn part1<R: BufRead>(reader: R) -> Result<usize> {
        let graph: Graph = Graph::from_str(reader.lines().flatten().collect::<Vec<_>>().join("\n"));
        // start DFS from "you" to "out"
        let mut memo: HashMap<(String, Vec<bool>), usize> = HashMap::new();
        let visited_required: Vec<bool> = Vec::new();
        Ok(graph.dfs_with_required(
            "you",
            "out",
            &vec![],
            visited_required,
            &mut memo,
        ))
    }

    assert_eq!(5, part1(BufReader::new(TEST1.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part1(input_file)?);
    println!("Result = {}", result);
    //endregion

    //region Part 2
    println!("\n=== Part 2 ===");

    fn part2<R: BufRead>(reader: R) -> Result<usize> {
        let graph: Graph = Graph::from_str(reader.lines().flatten().collect::<Vec<_>>().join("\n"));
        // start DFS from "svr" to "out", making sure that "dac" and "fft" are visited at least once
        let mut memo: HashMap<(String, Vec<bool>), usize> = HashMap::new();
        Ok(graph.dfs_with_required(
            "svr",
            "out",
            &vec!["dac", "fft"],
            vec![false; 2],
            &mut memo,
        ))
    }

    assert_eq!(2, part2(BufReader::new(TEST2.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part2(input_file)?);
    println!("Result = {}", result);
    //endregion

    Ok(())
}
