import collections
import time
import itertools

inputs = ['test1', 'input']

for input in inputs:

    for dimension in [3, 4]:
        with open(f'17/{input}', 'r') as f:
            active = set([(i, j) + (0,) * (dimension-2)for i, row in enumerate(f.read().split('\n'))
                          for j, cell in enumerate(row) if cell == '#'])

        # part 1

        start_time = time.time()

        neighbours = list(itertools.product([-1, 0, 1], repeat=dimension))

        for _ in range(6):
            neighbour_counts = collections.defaultdict(int)
            # for each active cube, add to the count of its neighbours
            for cube in active:
                for offset in neighbours:
                    if any(offset):
                        neighbour_counts[tuple(
                            a + b for a, b in zip(cube, offset))] += 1

            new_active = set()
            # loop through active neighbour counts of each cube
            for cube, count in neighbour_counts.items():
                if ((cube in active and 2 <= count <= 3)
                        or (cube not in active and count == 3)):
                    new_active.add(cube)
            active = new_active
        print(len(active))

        print(
            f'--- Time for dimension {dimension} for {input} is {(time.time() - start_time)} seconds ---')
