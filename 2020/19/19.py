import re


def regex_rule(rules, num, part2):
    rule = rules[num]
    # special rules to handle part 2
    if part2:
        if num == 8:
            return regex_rule(rules, 42, part2) + '+'
        elif num == 11:
            # hardcoded number of repeat to 9
            return '(' + '|'.join(
                f'{regex_rule(rules, 42, part2)}{{{i}}}{regex_rule(rules, 31, part2)}{{{i}}}'
                for i in range(1, 10)
            ) + ')'
    if '"' in rule:
        return rule.replace('"', '')
    else:
        res = [''.join(regex_rule(rules, int(subrule), part2)
                       for subrule in group.split(' '))
               for group in rule.split(' | ')]
        return '(' + '|'.join(res) + ')'


rules, messages = open(f'19/input', 'r').read().split('\n\n')

rules = {int(rule[0]): rule[1]
         for rule in (rule.split(': ') for rule in rules.split('\n'))}

messages = messages.split('\n')

# part 1
regex = regex_rule(rules, 0, part2=False)
print(sum((bool(re.fullmatch(regex, message))) for message in messages))

# part 2
rules[8] = '42 | 42 8'  # this is just rule(42)+
rules[11] = '42 31 | 42 11 31'  # this is (rule(42)rule(31))+

regex = regex_rule(rules, 0, part2=True)
print(sum((bool(re.fullmatch(regex, message))) for message in messages))
