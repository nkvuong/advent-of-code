import re

with open('02/input', 'r') as f:
    Lines = f.read().splitlines()

valid_passwords = 0

for line in Lines:
    lower, upper, char, temp, password = re.split('-|:| ', line)
    lower = int(lower)
    upper = int(upper)
    count = password.count(char)
    if (count >= lower) and (count <= upper):
        valid_passwords += 1
print(valid_passwords)

valid_passwords = 0

for line in Lines:
    pos1, pos2, char, temp, password = re.split('-|:| ', line)
    pos1 = int(pos1)
    pos2 = int(pos2)
    pos1_check = 0
    pos2_check = 0
    if pos1 < len(password)+1:
        pos1_check = password[pos1-1] == char
    if pos2 < len(password)+1:
        pos2_check = password[pos2-1] == char
    if (pos1_check + pos2_check) == 1:
        valid_passwords += 1
print(valid_passwords)
