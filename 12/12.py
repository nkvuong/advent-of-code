from os import error


def rotate(point, value):
    if value % 90 > 0:
        raise ValueError("rotation must be multiplier of 90")

    multiplier = [[1, 1], [1, -1], [-1, -1], [-1, 1]]

    if point[0] > 0:
        if point[1] > 0:
            quadrant = 0
        else:
            quadrant = 1
    else:
        if point[1] < 0:
            quadrant = 2
        else:
            quadrant = 3

    if value % 180 > 0:
        temp = point[0]
        point[0] = point[1]
        point[1] = temp

    quadrant = (quadrant + int(value/90)) % 4

    point[0] = multiplier[quadrant][0] * abs(point[0])
    point[1] = multiplier[quadrant][1] * abs(point[1])


inputs = ['test1', 'input']
NORTH, EAST, SOUTH, WEST = 0, 1, 2, 3
directions = ['N', 'E', 'S', 'W']
movement = {
    'N': (1, 1),
    'E': (0, 1),
    'S': (1, -1),
    'W': (0, -1)
}

for input in inputs:

    # Part 1
    with open(f'12/{input}', 'r') as f:
        instructions = [(r[0], int(r[1:])) for r in f.read().split("\n")]

    heading = EAST
    location = [0, 0]

    for action, value in instructions:
        if action == 'F':
            action = directions[heading]
        if action in 'NEWS':
            location[movement[action][0]] += movement[action][1] * value
        elif action in 'LR':
            heading = (heading + int(value/90) *
                       (-1 if action == 'L' else 1)) % 4
        # print(f'Position {location} heading {directions[heading]}')

    print(f'Part 1: {sum(map(abs, location))}')

    # Part 2
    location = [0, 0]
    waypoint = [10, 1]

    for action, value in instructions:
        if action == 'L':
            action = 'R'
            value = 360-value
        if action == 'F':
            location = [x + value * y for x, y in zip(location, waypoint)]
        if action in 'NEWS':
            waypoint[movement[action][0]] += movement[action][1] * value
        elif action == 'R':
            rotate(waypoint, value)

        # print(f'Position {location} waypoint {waypoint}')

    print(f'Part 2: {sum(map(abs, location))}')
