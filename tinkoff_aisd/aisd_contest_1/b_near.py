def min_diff(a, k):
    l, r = 0, len(a) - 1
    cur_res = float("inf")
    while l <= r:
        m = (l + r) // 2
        if abs(a[m] - k) < abs(cur_res - k):
            cur_res = a[m]
        elif abs(a[m] - k) == abs(cur_res - k):
            cur_res = min(a[m], cur_res)

        if a[m] < k:
            l = m + 1
        else:
            r = m - 1
    return cur_res


n, k = map(int, input().split())
seq = [int(x) for x in input().split()]
q_seq = [int(x) for x in input().split()]

for question in q_seq:
    print(min_diff(seq, question))
