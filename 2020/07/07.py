import re


def can_contain_gold(colour, rule_set):
    if len(rule_set[colour]) == 0:
        return False
    elif 'shiny gold' in rule_set[colour].keys():
        return True
    else:
        for bag in rule_set[colour].keys():
            if can_contain_gold(bag, rule_set):
                return True
    return False


def number_of_bags(colour, rule_set):
    total = 0
    if len(rule_set[colour]) == 0:
        total = 0
    else:
        for bag in rule_set[colour].keys():
            total += rule_set[colour][bag] * \
                (number_of_bags(bag, rule_set) + 1)
    return total


with open('07/input', 'r') as f:
    rules_input = f.read().split("\n")

rules = dict()

# parse the rules
for rule in rules_input:
    bags = re.split(' bags contain | bags?, | bags?.', rule)
    allowed_colours = dict()
    for bag in bags[1:-1]:
        if bag != 'no other':
            allowed_colours[bag[2:]] = int(bag[:1])
    rules[bags[0]] = allowed_colours

# loop through the ruleset to populate bag_contain_gold
available_colours = sum(can_contain_gold(colour, rules)
                        for colour in rules.keys())

gold_subbags = number_of_bags('shiny gold', rules)

print(
    f'Number of bags colours that can hold a gold bag is {available_colours}')

print(
    f'A gold bag can hold {gold_subbags} bags')
