def dfs(missed, visited, graph, node):
    if node not in graph.keys():
        missed.add(node)
        visited.add(node)
    if node not in visited:
        visited.add(node)
        for neighbour in graph[node]:
            dfs(missed, visited, graph, neighbour)


with open('08/input', 'r') as f:
    instructions = f.read().split("\n")

acc = 0
line = 0
visited = set()

while True:
    if line in visited:
        break
    op, arg = instructions[line][:3], instructions[line][4:]
    visited.add(line)
    line += int(arg) if op == 'jmp' else 1
    acc += int(arg) if op == 'acc' else 0

print(f'Accumulator before infinite loop is {acc}')

targets = dict()

for index, instruction in enumerate(instructions):
    op, arg = instruction[:3], instruction[4:]
    target = index - 1 + int(arg) if op == 'jmp' else index
    if target in targets:
        targets[target].append(index - 1)
    else:
        targets[target] = [index-1]

visited = set()
missed = set()
dfs(missed, visited, targets, len(instructions) - 1)

for candidate in missed:

    instructions[candidate] = instructions[candidate].replace(
        'jmp', 'nop') if 'jmp' in instructions[candidate] else instructions[candidate].replace('nop', 'jmp')

    acc = 0
    line = 0
    visited = set()
    while True:
        if line in visited:
            break
        if line >= len(instructions):
            print(
                f'Bad instruction is line {candidate+1}: {instructions[candidate]}. Final accumulator is {acc}')
            break
        op, arg = instructions[line][:3], instructions[line][4:]
        visited.add(line)
        line += int(arg) if op == 'jmp' else 1
        acc += int(arg) if op == 'acc' else 0

    instructions[candidate] = instructions[candidate].replace(
        'jmp', 'nop') if 'jmp' in instructions[candidate] else instructions[candidate].replace('nop', 'jmp')
