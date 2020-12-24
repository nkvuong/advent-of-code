import collections
from copy import deepcopy

tiles_input = open(f'24/input', 'r').read().split('\n')

# using hexagon axial coordinates (https://www.redblobgames.com/grids/hexagons/)

movement = {
    'e': (1, 0),
    'w': (-1, 0),
    'nw': (0, -1),
    'se': (0, 1),
    'ne': (1, -1),
    'sw': (-1, 1)
}

# part 1

black_tiles = collections.defaultdict(int)

for tile in tiles_input:
    coord = (0, 0)
    pos = 0
    dir = ''
    while len(tile) > 0:
        if tile.startswith('e') or tile.startswith('w'):
            dir = tile[:1]
            tile = tile[1:]
        else:
            dir = tile[:2]
            tile = tile[2:]
        coord = tuple(map(sum, zip(coord, movement[dir])))
    black_tiles[coord] += 1

black_tiles = {tile for tile, count in black_tiles.items() if count % 2 == 1}

print(len(black_tiles))

# part 2

for _ in range(100):
    neighbours = collections.defaultdict(int)
    new_black_tiles = deepcopy(black_tiles)
    # set count of black neighbours for each black tile to 0
    for tile in black_tiles:
        neighbours[tile] = 0
    # for each black tile, increase the black count of its neighbours
    for tile in black_tiles:
        for dir in movement.values():
            neighbours[tuple(map(sum, zip(tile, dir)))] += 1
    # loop through all the neighbours and flip accordingly
    for tile in neighbours:
        if tile in black_tiles:
            if neighbours[tile] == 0 or neighbours[tile] > 2:
                new_black_tiles.remove(tile)
        elif neighbours[tile] == 2:
            new_black_tiles.add(tile)
    black_tiles = new_black_tiles

print(len(black_tiles))
