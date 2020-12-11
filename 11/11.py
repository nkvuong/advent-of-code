def occupied_directions(row, col, seat_layout, limit):
    max_row, max_col = len(seat_layout), len(seat_layout[0])
    occupied_count = 0
    for drow, dcol in adjacency:
        for dist in range(1, limit+1):
            # boundaries check
            if row+dist*drow in range(max_row) and col+dist*dcol in range(max_col):
                if seat_layout[row+dist*drow][col+dist*dcol] == OCCUPIED:
                    occupied_count += 1
                    break
                elif seat_layout[row+dist*drow][col+dist*dcol] == EMPTY:
                    break

    return occupied_count


inputs = ['test1', 'input']  # the adjacency matrix
adjacency = [(i, j) for i in (-1, 0, 1)
             for j in (-1, 0, 1) if not (i == j == 0)]
OCCUPIED, EMPTY, FLOOR = '#L.'

for input in inputs:

    # Part 1
    with open(f'11/{input}', 'r') as f:
        original = [list(map(''.join, zip(*[iter(r)])))
                    for r in f.read().split("\n")]

    from copy import deepcopy
    seat_layout = deepcopy(original)
    changed = True
    while changed:

        previous = deepcopy(seat_layout)
        changed = False
        for r, row in enumerate(seat_layout):
            for c, cell in enumerate(row):
                if cell == FLOOR:
                    continue
                if cell == EMPTY and occupied_directions(r, c, previous, 1) == 0:
                    seat_layout[r][c] = OCCUPIED
                    changed = True
                if cell == OCCUPIED and occupied_directions(r, c, previous, 1) >= 4:
                    seat_layout[r][c] = EMPTY
                    changed = True

    print(f'Part 1: {sum(r.count(OCCUPIED) for r in seat_layout)}')

    # Part 2 - re-read the input
    seat_layout = deepcopy(original)
    max_dist = len(seat_layout)
    changed = True
    while changed:
        previous = deepcopy(seat_layout)
        changed = False
        for r, row in enumerate(seat_layout):
            for c, cell in enumerate(row):
                if cell == FLOOR:
                    continue
                if cell == EMPTY and occupied_directions(r, c, previous, max_dist) == 0:
                    seat_layout[r][c] = OCCUPIED
                    changed = True
                if cell == OCCUPIED and occupied_directions(r, c, previous, max_dist) >= 5:
                    seat_layout[r][c] = EMPTY
                    changed = True

    print(f'Part 2: {sum(r.count(OCCUPIED) for r in seat_layout)}')
