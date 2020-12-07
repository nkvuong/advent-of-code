import re

with open('06/input', 'r') as f:
    forms = f.read().split("\n")

# add blank row at the end

forms.append('')
total_any_yes = 0
total_all_yes = 0

# remove duplicate answers using a set
any_group_answers = set()

# check for all yes answers using a dict, total yes = total answer
all_group_answers = {'total': 0}

for answer in forms:
    if answer == '':
        total_any_yes += len(any_group_answers)
        any_group_answers = set()

        for q in all_group_answers:
            if (all_group_answers[q] == all_group_answers['total']) and (q != 'total'):
                total_all_yes += 1
        all_group_answers = {'total': 0}

    else:
        for q in answer:
            any_group_answers.add(q)
            if q in all_group_answers:
                all_group_answers[q] += 1
            else:
                all_group_answers[q] = 1
        all_group_answers['total'] += 1

print(f'Qs with any yes: {total_any_yes}')
print(f'Qs with all yes: {total_all_yes}')
