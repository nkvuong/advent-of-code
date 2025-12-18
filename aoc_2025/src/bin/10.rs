use anyhow::*;
use aoc_2025::*;
use code_timing_macros::time_snippet;
use const_format::concatcp;
use std::collections::HashMap;
use std::fs::File;
use std::io::{BufRead, BufReader};
use std::usize::MAX;
use std::hash::{Hash, DefaultHasher, Hasher};

const DAY: &str = "10";
const INPUT_FILE: &str = concatcp!("input/", DAY, ".txt");

const TEST: &str = "\
[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
";

struct Machine {
    lights: Vec<usize>,
    buttons: Vec<Vec<usize>>,
    joltages: Vec<usize>,
}

impl Machine {
    fn from_str(s: String) -> Self {
        let parts: Vec<&str> = s.split(' ').collect();
        let lights_str = parts[0].trim_matches(&['[', ']'] as &[_]);
        let lights: Vec<usize> = lights_str
            .chars()
            .map(|c| if c == '#' { 1 } else { 0 })
            .collect();
        let buttons: Vec<Vec<usize>> = parts[1..parts.len() - 1]
            .iter()
            .map(|btn_str| {
                let btn_clean = btn_str.trim_matches(&['(', ')'] as &[_]);
                btn_clean
                    .split(',')
                    .map(|num_str| num_str.parse::<usize>().unwrap())
                    .collect()
            })
            .collect();

        let joltage_str = parts.last().unwrap().trim_matches(&['{', '}'] as &[_]);
        let joltages: Vec<usize> = joltage_str
            .split(',')
            .map(|num_str| num_str.parse::<usize>().unwrap())
            .collect();

        Machine {
            lights,
            buttons,
            joltages,
        }
    }

    fn get_combinations(
        buttons: &Vec<Vec<usize>>,
        num_lights: usize,
    ) -> HashMap<Vec<usize>, Vec<usize>> {
        let num_buttons = buttons.len();
        let mut combinations = HashMap::with_capacity(1 << num_buttons);
        for combo in 0..(1 << num_buttons) {
            let mut lights = vec![0; num_lights];
            let mut current = Vec::with_capacity(num_buttons);
            for i in 0..num_buttons {
                if (combo & (1 << i)) != 0 {
                    current.push(i);
                    for &light_index in &buttons[i] {
                        lights[light_index] += 1;
                    }
                }
            }
            combinations.insert(current.clone(), lights.clone());
        }
        combinations
    }

    fn min_light_pressed(&self) -> usize {
        // create all possible combination of button presses
        // find the minimum number of presses that turns off all lights
        let mut min_presses = MAX;
        for (combo, presses) in Machine::get_combinations(&self.buttons, self.lights.len()) {
            if self
                .lights
                .iter()
                .zip(presses)
                .all(|(&state, count)| count % 2 == state)
            {
                min_presses = min_presses.min(combo.len());
            }
        }
        min_presses
    }

    fn min_joltage_pressed(&mut self) -> usize {
        let mut cache: HashMap<u64, usize> = HashMap::new();
        let combinations = Machine::get_combinations(&self.buttons, self.lights.len());
        let mut sorted_combos: Vec<_> = combinations.into_iter().collect();
        sorted_combos.sort_by_key(|(combo, _)| combo.len());
        Machine::min_joltage_pressed_helper(
            &sorted_combos,
            &self.joltages,
            &mut cache,
        )
    }

    fn min_joltage_pressed_helper(
        combinations: &Vec<(Vec<usize>, Vec<usize>)>,
        joltages: &Vec<usize>,
        cache: &mut HashMap<u64, usize>,
    ) -> usize {
        // base case: all joltages are zero
        if joltages.iter().all(|&j| j == 0) {
            return 0;
        }

        // check cache first
        let mut hash = DefaultHasher::new();
        joltages.hash(&mut hash);
        if let Some(&result) = cache.get(&hash.finish()) {
            return result;
        }

        let mut min_presses = MAX;

        // find the possible button combinations that can reach the current joltage state
        for (combo, presses) in combinations {
            // early pruning: if current combo + minimum possible future is worse, stop
            if combo.len() >= min_presses {
                break;
            }

            // check if this combination is valid for current joltages
            if !joltages
                .iter()
                .zip(presses.iter())
                .all(|(&joltage, &press)| (joltage % 2 == press % 2) && joltage >= press)
            {
                continue;
            }
            // recurse with reduced joltages
            let new_joltages: Vec<usize> = joltages
                .iter()
                .zip(presses.iter())
                .map(|(&joltage, &press)| (joltage - press) / 2)
                .collect();

            let new_presses =
                Machine::min_joltage_pressed_helper(combinations, &new_joltages, cache);

            if new_presses != MAX {
                min_presses = min_presses.min(combo.len() + 2 * new_presses);
            }
        }

        cache.insert(hash.finish(), min_presses);
        min_presses
    }
}

fn main() -> Result<()> {
    start_day(DAY);

    //region Part 1
    println!("=== Part 1 ===");

    fn part1<R: BufRead>(reader: R) -> Result<usize> {
        let machines: Vec<Machine> = reader.lines().flatten().map(Machine::from_str).collect();
        Ok(machines.into_iter().map(|m| m.min_light_pressed()).sum())
    }

    assert_eq!(7, part1(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part1(input_file)?);
    println!("Result = {}", result);
    //endregion

    //region Part 2
    println!("\n=== Part 2 ===");

    fn part2<R: BufRead>(reader: R) -> Result<usize> {
        let machines: Vec<Machine> = reader.lines().flatten().map(Machine::from_str).collect();
        Ok(machines
            .into_iter()
            .map(|mut m| m.min_joltage_pressed())
            .sum())
    }

    assert_eq!(33, part2(BufReader::new(TEST.as_bytes()))?);

    let input_file = BufReader::new(File::open(INPUT_FILE)?);
    let result = time_snippet!(part2(input_file)?);
    println!("Result = {}", result);
    //endregion

    Ok(())
}
