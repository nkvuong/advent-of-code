import math

# Function to calculate k for given a, b, m such that a^k = b (mod m), using baby step giant step


def discrete_log(a, b, m):
    n = int(math.sqrt(m) + 1)

    value = [0] * m

    # Store all values of a^(n*i) of LHS
    for i in range(n, 0, -1):
        value[pow(a, i * n, m)] = i

    for j in range(n):

        # Calculate (a ^ j) * b and check
        # for collision
        cur = (pow(a, j, m) * b) % m

        # If collision occurs i.e., LHS = RHS
        if (value[cur]):
            ans = value[cur] * n - j

            # Check whether ans lies below m or not
            if (ans < m):
                return ans

    return -1


door_pub = 1614360
card_pub = 7734663
MOD = 20201227

door_loop = discrete_log(7, door_pub, MOD)
card_loop = discrete_log(7, card_pub, MOD)

print(pow(door_pub, card_loop, 20201227))
