from functools import reduce

# chinese remainder theorem - solve for
# a mod n1 = a1
# a mod n2 = a2
# ...
# a mod nx = ax


def chinese_remainder(n, a):
    s = 0
    prod = reduce(lambda a, b: a*b, n)
    for n_i, a_i in zip(n, a):
        p = prod // n_i
        s += a_i * mul_inv(p, n_i) * p
    return s % prod

# calculate modular inverse using extended euclidean algo


def mul_inv(a, b):
    b0 = b
    x0, x1 = 0, 1
    if b == 1:
        return 1
    while a > 1:
        q = a // b
        a, b = b, a % b
        x0, x1 = x1 - q * x0, x0
    if x1 < 0:
        x1 += b0
    return x1


inputs = ['test1', 'input']
MAX_WAIT = 10000
for input in inputs:

    # Part 1
    with open(f'13/{input}', 'r') as f:
        timestamp, timetable = f.read().split("\n")
        timestamp = int(timestamp)
        timetable = [int(bus) if bus !=
                     'x' else 0 for bus in timetable.split(',')]
        waiting_time = min([
            (bus,
             bus - timestamp % bus if bus > 0 else MAX_WAIT)
            for bus in timetable
        ], key=lambda t: t[1])
        print(waiting_time[0]*waiting_time[1])

    # Part 2
    remainder = [bus - (i % bus) if bus > 0 else -
                 1 for i, bus in enumerate(timetable)]
    # remove unnecessary x
    timetable = [i for i in timetable if i > 0]
    remainder = [i for i in remainder if i > -1]
    # good thing all the buses are coprimes already
    print(chinese_remainder(timetable, remainder))
