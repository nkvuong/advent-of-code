import re
from itertools import compress


def is_valid(field, rule):
    l_1, u_1 = (int(d) for d in rule[1].split('-'))
    l_2, u_2 = (int(d) for d in rule[2].split('-'))
    if (l_1 <= int(field) <= u_1) or (l_2 <= int(field) <= u_2):
        return True
    else:
        return False


inputs = ['test1', 'test2', 'input']
for input in inputs:

    with open(f'16/{input}', 'r') as f:
        rules, my_ticket, nearby = [i.split('\n') for i in
                                    re.split('\n\nyour ticket:\n|\n\nnearby tickets:\n', f.read())]
    rules = [re.split(':| or ', rule) for rule in rules]
    nearby = [list(map(int, ticket.split(','))) for ticket in nearby]
    my_ticket = list(map(int, my_ticket[0].split(',')))

    error_rate = 0
    valid_ticket = [True for _ in nearby]

    for i, ticket in enumerate(nearby):
        for field in ticket:
            valid = any([is_valid(field, rule) for rule in rules])
            error_rate += field * valid
            valid_ticket[i] = valid_ticket[i] and valid

    # filter out invalid tickets
    nearby = list(compress(nearby, valid_ticket))
    # part 1
    print(error_rate)

    # part 2
    field_mapping = [
        # intersect potential field for each each element of the list, e.g. 1st field of each ticket, 2nd field & so on
        set.intersection(*x) for x in zip(
            # loop through each ticket to create a list of potential match for each field
            *[
                # create a set of potential match for each field in the ticket
                [
                    set([rule[0] for rule in rules if is_valid(field, rule)])
                    for field in ticket
                ]
                for ticket in nearby
            ]
        )
    ]

    # keep removing found fields from the the other potential fields
    while sum(map(len, field_mapping)) > len(field_mapping):
        found = set.union(*[f for f in field_mapping if len(f) == 1])
        field_mapping = [f.difference(found) if len(f) > 1 else f
                         for f in field_mapping]
    print(field_mapping)

    value = 1
    for i, field in enumerate(field_mapping):
        (name,) = field
        if 'departure' in name:
            value *= my_ticket[i]

    print(value)
