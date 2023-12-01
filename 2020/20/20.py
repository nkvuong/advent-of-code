import re
import itertools
import collections
import math
from functools import reduce
from copy import deepcopy


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


def rotflip(tile, option):
    # rotate right
    if option == R:
        return rotate(tile)
    # rotate left
    if option == L:
        return rotate(rotate(rotate(tile)))
    # rotate 180
    if option == R180:
        return rotate(rotate(tile))
    # rotate right then flipped
    if option == RF:
        return flip(rotate(tile))
    # rotate left then flipped
    if option == LF:
        return flip(rotate(rotate(rotate(tile))))
    # just flip vertical
    if option == FV:
        return flip(tile)
    # rotate 180 then flip (flip horizontal)
    if option == FH:
        return flip(rotate(rotate(tile)))
    # do nothing
    return tile


def border(tiles, num, direction, option):

    tile = rotflip(tiles[num], option)

    result = ''
    if direction == N:
        result = ''.join(tile[0])
    elif direction == E:
        result = ''.join([row[-1] for row in tile])
    elif direction == S:
        result = ''.join(tile[-1])
    elif direction == W:
        result = ''.join([row[0] for row in tile])

    return result


N, E, S, W = 0, 1, 2, 3

# 8 options = do nothing, 3 rotations, flip only, 3 flipped rotations
R, L, R180, FV, RF, LF, FH = 1, 2, 3, 4, 5, 6, 7

DIRECTIONS = [(-1, 0), (0, 1), (1, 0), (0, -1)]

TILE_SIZE = 10
COMBINATION_NUM = 8
DIRECTION_NUM = 4

input = open(f'20/input', 'r').read().split('\n\n')

neighbours = collections.defaultdict(set)

tiles = {int(tile[1]): tile[2].split('\n')
         for tile in (re.split('Tile |:\n', t) for t in input)}

TILE_IMAGE = int(math.sqrt(len(tiles)))

# part 1

# create a list of 8 edges for each tile (one in each direction, and the reversed one)

edges = {}
for num, tile in tiles.items():
    temp = [border(tiles, num, direction, 0)
            for direction in range(DIRECTION_NUM)]
    edges[num] = set([x for x in temp] + [x[::-1] for x in temp])

# compare each tile and create a list of neighbours

for each in itertools.permutations(edges, 2):
    if edges[each[0]] & edges[each[1]]:  # two tile has common edges
        neighbours[each[0]].add(each[1])

# tiles with 2 neighbours are the corners
print(reduce((lambda x, y: x * y),
             [neighbour for neighbour in neighbours if len(neighbours[neighbour]) == 2]))

# part 2 - now need to join the tiles to produce the image

loc = [[None for _ in range(TILE_IMAGE)] for _ in range(TILE_IMAGE)]

# set a corner, and the next two tiles
corner = [neighbour for neighbour in neighbours if len(
    neighbours[neighbour]) == 2][0]

used = set([corner]) | neighbours[corner]
loc[0][0] = corner
loc[0][1], loc[1][0] = neighbours[corner]

# find out which tiles fit where

while len(used) < len(tiles):
    for r in range(TILE_IMAGE):
        for c in range(TILE_IMAGE):
            if loc[r][c] is not None:
                continue
            # all unused tiles are possible
            options = set(
                [neighbour for neighbour in neighbours.keys() if neighbour not in used])
            for dr, dc in DIRECTIONS:
                # if the neighbour is already decided, used the neighbours list to narrow
                if r+dr in range(TILE_IMAGE) and c+dc in range(TILE_IMAGE) and loc[r+dr][c+dc]:
                    options &= neighbours[loc[r+dr][c+dc]]

            # assume there is only 1 tile for each position
            if len(options) == 1:
                chosen = options.pop()
                loc[r][c] = chosen
                used.add(chosen)

# brute force the direction of each tile

pieces = [[None for _ in range(TILE_IMAGE)] for _ in range(TILE_IMAGE)]

for r in range(TILE_IMAGE):
    for c in range(TILE_IMAGE):
        options = set(range(COMBINATION_NUM))
        for dr, dc in DIRECTIONS:
            option_dir = set()
            if r+dr in range(TILE_IMAGE) and c+dc in range(TILE_IMAGE):
                for rotflip_1 in range(COMBINATION_NUM):
                    for rotflip_2 in range(COMBINATION_NUM):
                        # match only the required neighbour
                        if dr == -1:  # up, so top & bottom
                            direction_1, direction_2 = N, S
                        elif dr == 1:  # down so bottom & top
                            direction_1, direction_2 = S, N
                        elif dc == 1:  # right so right & left
                            direction_1, direction_2 = E, W
                        else:  # left so left & right
                            direction_1, direction_2 = W, E
                        # add this rotation/flip as an option
                        if border(tiles, loc[r][c], direction_1, rotflip_1) == border(tiles, loc[r+dr][c+dc], direction_2, rotflip_2):
                            option_dir.add(rotflip_1)
                # the valid option is the one that aligns against all neighbours
                options &= option_dir
        if len(options) == 1:
            chosen = options.pop()
            pieces[r][c] = rotflip(tiles[loc[r][c]], chosen)

image = [[' ' for _ in range(TILE_IMAGE * (TILE_SIZE - 2))]
         for _ in range(TILE_IMAGE * (TILE_SIZE - 2))]

for r in range(TILE_IMAGE):
    for c in range(TILE_IMAGE):
        tile = pieces[r][c]
        for rr in range(1, len(tile)-1):
            for cc in range(1, len(tile[rr])-1):
                image[r*(TILE_SIZE-2)+(rr-1)][c *
                                              (TILE_SIZE-2)+(cc-1)] = tile[rr][cc]

# print('\n'.join([''.join(row) for row in image]))

MONSTER = ['                  # ',
           '#    ##    ##    ###',
           ' #  #  #  #  #  #   ']
MONSTER_LENGTH = len(MONSTER[0])
MONSTER_HEIGHT = len(MONSTER)

IMAGE_SIZE = len(image)

for option in range(6):
    has_monster = False
    edited_image = rotflip(image, option)
    is_monster = [[False for _ in range(IMAGE_SIZE)]
                  for _ in range(IMAGE_SIZE)]
    for r in range(IMAGE_SIZE):
        for c in range(IMAGE_SIZE):
            monster = True
            for mr in range(MONSTER_HEIGHT):
                for mc in range(MONSTER_LENGTH):
                    # monster would fall out of range
                    if not (r+mr in range(IMAGE_SIZE) and c+mc in range(IMAGE_SIZE)):
                        monster = False
                    else:
                        # monster cannot exist
                        if MONSTER[mr][mc] == '#' and edited_image[r+mr][c+mc] != '#':
                            monster = False
            if monster:  # monster exists in current location
                has_monster = True
                for mr in range(MONSTER_HEIGHT):
                    for mc in range(MONSTER_LENGTH):
                        if MONSTER[mr][mc] == '#':
                            is_monster[r+mr][c+mc] = True

    # Only one orientation has sea monsters
    if has_monster:
        #print('\n'.join([''.join(row) for row in edited_image]))
        print(sum([i == '#' and not m for im, mon in zip(
            edited_image, is_monster) for i, m in zip(im, mon)]))
