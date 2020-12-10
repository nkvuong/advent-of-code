def count_arrangements(adapter, distinct_arrangements):
    if distinct_arrangements[adapter] > 0:
        return distinct_arrangements[adapter]
    arrangement = 0
    for i in range(adapter-1, adapter-4, -1):
        if i in distinct_arrangements:
            arrangement += count_arrangements(i, distinct_arrangements)
    distinct_arrangements[adapter] = arrangement
    return arrangement


inputs = ['test1', 'test2', 'input']

for input in inputs:

    with open(f'10/{input}', 'r') as f:
        adapters = [int(i) for i in f.read().split("\n")]

    adapters.sort()

    current_joltage = 0
    jolt = [0, 0, 1]
    for adapter in adapters:
        jolt[adapter - current_joltage - 1] += 1
        current_joltage = adapter
    print(f'{jolt[0]} {jolt[2]} = {jolt[0]*jolt[2]}')

    distinct_arrangements = {0: 1}

    for adapter in adapters:
        distinct_arrangements[adapter] = 0

    print(count_arrangements(adapters[len(adapters)-1], distinct_arrangements))
