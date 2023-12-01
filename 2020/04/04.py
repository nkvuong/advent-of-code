import re


def is_between(value, low, high):
    try:
        int(value)
        return (int(value) >= low and int(value) <= high)
    except ValueError:
        return False


def match_rgx(rgx, value):
    m = rgx.match(value)
    if m:
        return True
    else:
        return False


with open('04/input', 'r') as f:
    passports = f.read().split("\n\n")

expected_fields = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']

valid_passports_1 = 0

for passport in passports:
    valid = True
    for expected_field in expected_fields:
        valid = valid and (expected_field in passport)
    valid_passports_1 += valid

print(valid_passports_1)

valid_passports_2 = 0

hcl_rgx = re.compile('^#[0-9a-f]{6}$', re.IGNORECASE)
ecl_rgx = re.compile('^amb|blu|brn|gry|grn|hzl|oth$', re.IGNORECASE)
pid_rgx = re.compile('^[0-9]{9}$')
hgt_rgx = re.compile('^[0-9]{2,3}(in|cm)$', re.IGNORECASE)

for passport in passports:
    fields = re.split('\n| ', passport)
    valid = True
    for expected_field in expected_fields:
        valid = valid and (expected_field in passport)
    for field in fields:
        key, value = field.split(':')
        if key == 'byr':
            valid = valid and is_between(value, 1920, 2002)
        elif key == 'iyr':
            valid = valid and is_between(value, 2010, 2020)
        elif key == 'eyr':
            valid = valid and is_between(value, 2020, 2030)
        elif key == 'hgt':
            if match_rgx(hgt_rgx, value):
                unit = value[-2:]
                value = value[:-2]
                if unit == 'in':
                    valid = valid and is_between(value, 59, 76)
                elif unit == 'cm':
                    valid = valid and is_between(value, 150, 193)
            else:
                valid = False
        elif key == 'hcl':
            valid = valid and match_rgx(hcl_rgx, value)
        elif key == 'ecl':
            valid = valid and match_rgx(ecl_rgx, value)
        elif key == 'pid':
            valid = valid and match_rgx(pid_rgx, value)
    valid_passports_2 += valid
print(valid_passports_2)
