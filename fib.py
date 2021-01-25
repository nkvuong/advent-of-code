fib = [0, 1]
for _ in range(50):
    fib.append(fib[-1] + fib[-2])

for i in range(10, 60, 10):
    print(fib[i])
    print(int(((1+5**.5)/2)**i/5**.5+.5))
