MOD = 10**9 + 7

class SegmentTree:
    def __init__(self, size):
        self.n = 1
        while self.n < size:
            self.n <<= 1
        self.size = size
        self.max_len = [0] * (2 * self.n)
        self.count = [0] * (2 * self.n)

    def update(self, idx, new_len, new_count):
        idx += self.n - 1  # Переход к листу
        if self.max_len[idx] < new_len:
            self.max_len[idx] = new_len
            self.count[idx] = new_count % MOD
        elif self.max_len[idx] == new_len:
            self.count[idx] = (self.count[idx] + new_count) % MOD
        else:
            return
        idx >>= 1
        while idx >= 1:
            left = 2 * idx
            right = 2 * idx + 1
            if self.max_len[left] > self.max_len[right]:
                self.max_len[idx] = self.max_len[left]
                self.count[idx] = self.count[left]
            elif self.max_len[right] > self.max_len[left]:
                self.max_len[idx] = self.max_len[right]
                self.count[idx] = self.count[right]
            else:
                self.max_len[idx] = self.max_len[left]
                self.count[idx] = (self.count[left] + self.count[right]) % MOD
            idx >>= 1

    def query(self, l, r):
        res_len = 0
        res_count = 0
        l += self.n - 1
        r += self.n - 1
        while l <= r:
            if l % 2 == 1:
                if self.max_len[l] > res_len:
                    res_len = self.max_len[l]
                    res_count = self.count[l]
                elif self.max_len[l] == res_len:
                    res_count = (res_count + self.count[l]) % MOD
                l += 1
            if r % 2 == 0:
                if self.max_len[r] > res_len:
                    res_len = self.max_len[r]
                    res_count = self.count[r]
                elif self.max_len[r] == res_len:
                    res_count = (res_count + self.count[r]) % MOD
                r -= 1
            l >>= 1
            r >>= 1
        return (res_len, res_count)

def main():
    import sys
    input = sys.stdin.read().split()
    n = int(input[0])
    a = list(map(int, input[1:n+1]))

    # Сжатие координат
    unique = sorted(a)
    rank = {v: i+1 for i, v in enumerate(sorted(set(a)))}
    m = len(rank)

    st = SegmentTree(m)

    for num in a:
        r = rank[num]
        if r == 1:
            current_max, current_cnt = 0, 0
        else:
            current_max, current_cnt = st.query(1, r-1)

        new_len = current_max + 1
        new_count = current_cnt if current_max != 0 else 1
        st.update(r, new_len, new_count)

    max_len, total = st.query(1, m)
    print(total % MOD)

if __name__ == "__main__":
    main()
