def execute(instructions):
    acc = 0
    line = 0
    visited = set()
    halting = False
    while True:
        if line in visited:
            break
        if line >= len(instructions):
            halting = True
            break
        op, arg = instructions[line][:3], int(instructions[line][4:])
        visited.add(line)
        line += arg if op == 'jmp' else 1
        acc += arg if op == 'acc' else 0
    return acc, halting, visited


def swap(op):
    if 'jmp' in op:
        return op.replace('jmp', 'nop')
    else:
        return op.replace('nop', 'jmp')


def swap_test(candidate, instructions):
    instructions[candidate] = swap(instructions[candidate])
    result = execute(instructions)
    instructions[candidate] = swap(instructions[candidate])
    return result


with open('08/input', 'r') as f:
    instructions = f.read().split("\n")

print(f'Accumulator before infinite loop is {execute(instructions)[0]}')

pointer = 0
visited = execute(instructions)[2]
potential_landing = [False for i in instructions]

# find all potential spaces that could lead to the end
for pointer in range(len(instructions)-1, 0, -1):
    if 'jmp -' in instructions[pointer]:
        break
    potential_landing[pointer] = True

candidate = 0
start = pointer

# last negative jump is already visited, so just need to change to nop to reach a potential space
if pointer in visited:
    candidate = pointer
else:
    while True:
        pointer -= 1
        op, arg = instructions[pointer][:3], int(instructions[pointer][4:])
        if potential_landing[pointer]:
            continue
        # a nop that could jump to to a potential space
        elif op == 'nop' and (pointer in visited) and potential_landing[pointer+arg]:
            candidate = pointer
            break
        # a jump that has not been visited, but would lead to a potential space
        elif op == 'jmp' and (pointer not in visited) and (not potential_landing[pointer]) and (potential_landing[pointer+arg]):
            # find the previous jump
            j = pointer - 1
            while not('jmp' in instructions[pointer]):
                j -= 1
            # if previous jump is already visited, change it to nop to be able to reach the next jump
            if j in visited:
                candidate = j
                break
            # if not visited, mark the space between these two jumps as potential
            else:
                potential_landing[j:pointer] = [True] * (pointer-j+1)
                pointer = start

print(
    f'Bad instruction is line {candidate+1}: {instructions[candidate]}. Final accumulator is {swap_test(candidate,instructions)[0]}')
