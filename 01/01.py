sum = 2020
f = open('01/input', 'r')
Lines = f.readlines()

Expenses = list()

for line in Lines:
    Expenses.append(int(line))

length = len(Expenses)
s = set()

for expense in Expenses:
    temp = sum - expense
    if (temp in s):
        print(f"{expense} {temp} {expense*temp}")
    s.add(expense)

for i in range(0, length-1):
    s = set()
    curr_sum = sum - Expenses[i]
    for j in range(i+1, length):
        temp = curr_sum - Expenses[j]
        if (temp in s):
            print(
                f"{Expenses[i]} {Expenses[j]} {temp} {Expenses[i]*Expenses[j]*temp}")
        s.add(Expenses[j])
