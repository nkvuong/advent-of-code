inputs = ['test1', 'input']
movement = {
    'N': 1,
    'E': 1j,
    'S': -1,
    'W': -1j,
    'R': 1j,
    'L': -1j
}

for input in inputs:

    # Part 1
    with open(f'12/{input}', 'r') as f:
        instructions = [(r[0], int(r[1:])) for r in f.read().split("\n")]

    heading = 0 + 1j
    location = 0 + 0j

    for action, value in instructions:
        if action == 'F':
            location += heading * value
        if action in 'NEWS':
            location += movement[action] * value
        elif action in 'LR':
            for _ in range((value // 90)):
                heading *= movement[action]
        # print(f'Position {location} heading {heading}')

    print(f'Part 1: {int(abs(location.real) + abs(location.imag))}')

    # Part 2
    location = 0 + 0j
    waypoint = 1 + 10j

    for action, value in instructions:
        if action == 'F':
            location += waypoint * value
        if action in 'NEWS':
            waypoint += movement[action] * value
        elif action in 'LR':
            for _ in range((value // 90)):
                waypoint *= movement[action]

        # print(f'Position {location} waypoint {waypoint}')

    print(f'Part 2: {int(abs(location.real) + abs(location.imag))}')
