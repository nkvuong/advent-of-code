from operator import add, mul

inputs = ['test1', 'input']
ops = {'+': add, '*': mul}
priority = {'+': 2, '*': 1}


def to_suffix(s, precendence):
    stack = []
    suffix = ''
    for val in s:
        if val in ops.keys():
            if precendence:
                # operator with lower priority cannot be on top of higher priority operator, so need to pop
                while stack and stack[-1] != '(' and priority[val] <= priority[stack[-1]]:
                    suffix += stack.pop()
            else:
                # ignore operator priority if precendence is false (part 1)
                while stack and stack[-1] != '(':
                    suffix += stack.pop()
            stack.append(val)
        elif val == '(':
            stack.append(val)
        # for ), need to pop the stack until the previous ( is reached
        elif val == ')':
            while stack[-1] != '(':
                suffix += stack.pop()
            # discard the ( as well
            stack.pop()
        else:
            suffix += val
    # output the rest of the stack to the suffix expression
    while stack:
        suffix += stack.pop()
    return suffix


def eval(s):
    stack = []
    for val in s:
        if val in ops.keys():
            n1 = stack.pop()
            n2 = stack.pop()
            stack.append(ops[val](n1, n2))
        else:
            stack.append(int(val))
    return stack.pop()


for input in inputs:

    with open(f'18/{input}', 'r') as f:
        homework = [expression.replace(' ', '')
                    for expression in f.read().split('\n')]

    # part 1
    print(sum(eval(to_suffix(expression, False)) for expression in homework))

    # part 2
    print(sum(eval(to_suffix(expression, True)) for expression in homework))
