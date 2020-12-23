def mixing_cup(cups, moves_num, max_cups):

    # set the pointers array for the list [1, 2, 3, 4, ..., max_cups]
    nex = [i + 1 for i in range(max_cups + 1)]

    # update the pointer for the first set of cups
    for i, label in enumerate(cups[:-1]):
        nex[label] = cups[i + 1]
    current_cup = cups[0]

    # set the the first cup pointer to the max cup
    if max_cups > len(cups):
        nex[-1] = current_cup
        nex[cups[-1]] = max(cups) + 1
    else:
        nex[cups[-1]] = current_cup

    for _ in range(moves_num):

        next_cup = nex[current_cup]
        next_cups = next_cup, nex[next_cup], nex[nex[next_cup]]

        # the next 3 cups will be removed, so move the pointer for the current cup
        nex[current_cup] = nex[nex[nex[next_cup]]]

        destination_cup = current_cup - 1 if current_cup > 1 else max_cups
        # keep subtracting one until the destination cup is not in the next 3
        while destination_cup in next_cups:
            destination_cup = destination_cup - 1 if destination_cup > 1 else max_cups
        # insert the 3 cups after destination by moving the pointer
        nex[nex[nex[next_cup]]] = nex[destination_cup]
        nex[destination_cup] = next_cup

        current_cup = nex[current_cup]

    return nex

# part 1


cups = list(map(int, '389125467'))
nex = mixing_cup(cups, 10, len(cups))
cup = nex[1]
res = str(cup)
while nex[cup] != 1:
    res += str(nex[cup])
    cup = nex[cup]
print(res)

nex = mixing_cup(cups, 100, len(cups))
cup = nex[1]
res = str(cup)
while nex[cup] != 1:
    res += str(nex[cup])
    cup = nex[cup]
print(res)

cups = list(map(int, '315679824'))
nex = mixing_cup(cups, 100, len(cups))
cup = nex[1]
res = str(cup)
while nex[cup] != 1:
    res += str(nex[cup])
    cup = nex[cup]
print(res)

# part 2

cups = list(map(int, '389125467'))
nex = mixing_cup(cups, 10000000, 1000000)
print(nex[1] * nex[nex[1]])

cups = list(map(int, '315679824'))
nex = mixing_cup(cups, 10000000, 1000000)
print(nex[1] * nex[nex[1]])
