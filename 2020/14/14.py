def floating_bit(value, start):
    # no X means only 1 value
    if not 'X' in value:
        return [''.join(value)]

    possible_values = []
    for position, bit in enumerate(value[start:], start):
        # need to branch if the bit is X
        if bit == 'X':
            # add both 0 & 1 to the list
            value[position] = '0'
            possible_values.extend(floating_bit(value, position))
            value[position] = '1'
            possible_values.extend(floating_bit(value, position))
            value[position] = 'X'
    return possible_values


def mask_v2(bitmask, value):
    # pad the binary value to 36 bit
    value_binary = bin(value)[2:].rjust(36, '0')
    result = [bit if mask == '0' else mask
              for bit, mask in zip(value_binary, bitmask)]
    return floating_bit(result, 0)


def mask_v1(bitmask, value):
    # pad the binary value to 36 bit
    value_binary = bin(value)[2:].rjust(36, '0')
    result = ''.join([bit if mask == 'X' else mask
                      for bit, mask in zip(value_binary, bitmask)])
    return int(result, 2)


inputs = ['test2', 'input']
for input in inputs:

    with open(f'14/{input}', 'r') as f:
        programs = [r.split(' = ') for r in f.read().split("\n")]

        # Part 1
        bitmask = ''
        memory = {}
        for cmd, param in programs:
            if cmd == 'mask':
                bitmask = param
            else:
                memory[int(cmd[4:-1])] = mask_v1(bitmask, int(param))
        print(sum(memory.values()))

        # Part 2
        memory = {}
        for cmd, param in programs:
            if cmd == 'mask':
                bitmask = param
            else:
                for value in mask_v2(bitmask, int(cmd[4:-1])):
                    memory[int(value, 2)] = int(param)
        print(sum(memory.values()))
