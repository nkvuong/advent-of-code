import re

def binary_search(lower,upper,seq):
    for element in seq:
        if element == 'F' or element == 'L':
            upper = (lower+upper)//2
        else: lower = (lower+upper)//2 + 1
    return lower

with open('05/input', 'r') as f:
    bps = f.read().split("\n")

highest_bp = 0
min_seat = 0
max_seat = 127 * 8 + 7
available_seats = set()

for i in range(min_seat,max_seat+1):
    available_seats.add(i)

for bp in bps:
    row = binary_search(0,127,bp[:7])
    column = binary_search(0,7,bp[7:])
    seat_id = row * 8 + column
    highest_bp = max(seat_id,highest_bp)
    available_seats.remove(seat_id)
    # print(f'{row} {column} {seat_id}')

for i in range(min_seat,max_seat):
    if not (i in available_seats):
        break
    else:
        available_seats.remove(i)

for i in range(max_seat,min_seat,-1):
    if not (i in available_seats):
        break
    else:
        available_seats.remove(i)

print(highest_bp)
print(available_seats)