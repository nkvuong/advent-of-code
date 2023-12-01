import functools
import operator

with open('03/input', 'r') as f:
    lines = f.read().splitlines()

tree_counts = []
slopes = [(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)]

for slope in slopes:
    tree_count = 0
    pos = 0
    slope_right = slope[0]
    slope_down = slope[1]

    for i, line in enumerate(lines):
        if i % slope_down == 0:
            length = len(line)
            landscape = [line[i:i + 1] for i in range(0, length, 1)]
            if (landscape[pos] == '#'):
                tree_count += 1
            pos = (pos + slope_right) % length
    tree_counts.append(tree_count)

result = functools.reduce(operator.mul, tree_counts)

print(tree_counts)
print(result)
