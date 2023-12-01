inputs = [
    [0, 3, 6],
    [1, 3, 2],
    [2, 1, 3],
    [1, 2, 3],
    [2, 3, 1],
    [3, 2, 1],
    [3, 1, 2],
    [7, 12, 1, 0, 16, 2]
]

for input in inputs:
    for part in [2020, 30000000]:
        spoken, seen = 0, {num: pos+1 for pos, num in enumerate(input)}

        for turn in range(len(input)+1, part):
            seen[spoken], spoken = turn, 0 if spoken not in seen else turn-seen[spoken]
        print(
            f'For starting number {input}, the {part}th number spoken is {spoken}')
