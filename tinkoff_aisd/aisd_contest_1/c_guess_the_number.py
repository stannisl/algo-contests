import sys


def query():
    sys.stdout.flush()
    return input()


n = int(input())

l, r = 0, n

while l < r:
    m = (l + r + 1) // 2
    print(m)
    c = query()
    if c == "<":
        r = m - 1
    if c == ">=":
        l = m

print(f"! {l}")
