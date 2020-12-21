import re
import collections
from copy import deepcopy

input = open(f'21/input', 'r').read().split('\n')

foods = {food[0]: re.split(', |\)', food[1])[:-1]
         for food in (t.split(' (contains ') for t in input)}

# part 1

# create a list of all ingredients
all_ingredients = set.union(*[set(food.split(' ')) for food in foods.keys()])


# for each food, add the ingredients to possible allergens set
possible_allergens = collections.defaultdict(
    lambda: deepcopy(all_ingredients))

for food, allergies in foods.items():
    for allergy in allergies:
        possible_allergens[allergy] &= set(food.split(' '))

# the rest of the food definitely does not contain an allergen
non_allergen = all_ingredients.difference(
    set.union(*possible_allergens.values()))

# count the number of times a non-allergen ingredient appears
ans = sum([ingredient in non_allergen for food in foods.keys()
           for ingredient in food.split(' ')])

print(ans)

# part 2

# the allergen with 1 possible ingredient is the right one, remove it from all other possible list
while (len(possible_allergens) < sum(len(x) for x in possible_allergens.values())):
    for allergy, ingredients in possible_allergens.items():
        if len(ingredients) == 1:
            for all, ing in possible_allergens.items():
                if all != allergy:
                    ing -= ingredients

# sort the dictionary and print the comma separated values

print(','.join([s[1].pop()
                for s in sorted(possible_allergens.items())]))
