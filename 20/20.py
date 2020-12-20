import re
import itertools
import collections
import math
from functools import reduce


def rotate(tile):
    row = len(tile)
    column = len(tile)
    # takes (r,c) to (c, R-1-r)
    rotate_tile = [[' ' for _ in range(row)] for _ in range(column)]
    for r in range(row):
        for c in range(column):
            rotate_tile[c][row-1-r] = tile[r][c]
    return rotate_tile


def flip(tile):
    return list(reversed(tile))


def flip2(tile):
    return [list(reversed(row)) for row in tile]


def border(tiles, num, direction, rotflip):
    result = ''

    if rotflip == R:
        direction = (direction + 1) % 4

    if rotflip == L:
        direction = (direction + 3) % 4

    if rotflip == R180:
        direction = (direction + 2) % 4

    if ((rotflip == FH) and (direction in (E, W))) or ((rotflip == FV) and (direction in (S, N))):
        direction = (direction + 2) % 4

    if direction == N:
        result = tiles[num][:10]
    elif direction == E:
        result = tiles[num][9::11]
    elif direction == S:
        result = tiles[num][-10:]
    else:
        result = tiles[num][::11]

    if ((rotflip == FH) and (direction in (S, N))) or ((rotflip == FV) and (direction in (E, W))):
        result = result[::-1]

    return result


N, E, S, W = 0, 1, 2, 3
R, L, R180, FH, FV = 1, 2, 3, 4, 5

input = open(f'20/test', 'r').read().split('\n\n')

borders = collections.defaultdict(set)

tiles = {int(tile[1]): tile[2]
         for tile in (re.split('Tile |:\n', t) for t in input)}

image_size = int(math.sqrt(len(tiles)))

# part 1

for each in itertools.permutations(tiles, 2):
    for direction_1 in range(4):
        for rotflip_1 in range(6):
            for direction_2 in range(4):
                for rotflip_2 in range(6):
                    if border(tiles, each[0], direction_1, rotflip_1) == border(tiles, each[1], direction_2, rotflip_2):
                        borders[each[0]].add(each[1])

print(reduce((lambda x, y: x * y),
             [border for border in borders if len(borders[border]) == 2]))

MONSTER = ['                  # ',
           '#    ##    ##    ###',
           ' #  #  #  #  #  #   ']

D = [(-1, 0), (0, 1), (1, 0), (0, -1)]

loc = [[None for _ in range(image_size)] for _ in range(image_size)]

# set a corner, and the next two tiles
corner = [border for border in borders if len(borders[border]) == 2][0]

used = set([corner]) | borders[corner]
loc[0][0] = corner
loc[0][1], loc[1][0] = borders[corner]

# find out which tiles fit where

while len(used) < len(tiles):
    for r in range(image_size):
        for c in range(image_size):
            if loc[r][c] is not None:
                continue
            options = set([k for k in borders.keys() if k not in used])
            for dr, dc in D:
                rr, cc = r+dr, c+dc
                if rr in range(image_size) and cc in range(image_size) and loc[rr][cc]:
                    options &= borders[loc[rr][cc]]
            if len(options) == 1:
                chosen = options.pop()
                loc[r][c] = chosen
                used.add(chosen)

# brute force the direction of each tile

pieces = [[None for _ in range(image_size)] for _ in range(image_size)]
