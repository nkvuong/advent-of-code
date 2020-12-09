def is_valid(message, preamble):
    for number in preamble:
        if (message-number) in preamble:
            return True
    return False


with open('09/input', 'r') as f:
    messages = [int(i) for i in f.read().split("\n")]

preamble = set()
preamble_start = 0
preamble_length = 25
target = 0

for i in range(0, preamble_length):
    preamble.add(messages[i])

for i in range(preamble_length, len(messages)):
    if not is_valid(messages[i], preamble):
        print(f'First invalid number is {messages[i]}')
        target = messages[i]
        break
    preamble.remove(messages[preamble_start])
    preamble_start += 1
    preamble.add(messages[i])

start = 0
end = -1
sum = 0

while (end <= len(messages)):
    if sum < target:
        end += 1
        sum += messages[end]
    elif sum > target:
        sum -= messages[start]
        start += 1
    else:
        break

min = target
max = 0
for i in range(start, end+1):
    if max < messages[i]:
        max = messages[i]
    if min > messages[i]:
        min = messages[i]

print(
    f'The encryption weakness is {min+max}, between {start}:{messages[start]} and {end}:{messages[end]}')
