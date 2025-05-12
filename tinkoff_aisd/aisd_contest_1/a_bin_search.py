def bin_search(seq, elem):
    l, r = 0, len(seq) - 1

    while l <= r:
        m = (l + r) // 2
        if seq[m] == elem:
            return True
        elif seq[m] < elem:
            l = m + 1
        elif seq[m] > elem:
            r = m - 1

    return False


n, k = map(int, input().split())
seq = [int(x) for x in input().split()]
q_seq = [int(x) for x in input().split()]

for question in q_seq:
    print("YES" if bin_search(seq, question) else "NO")
