# https://coderun.yandex.ru/selections/2024-summer-backend/problems/splitting-into-terms?compiler=python

import sys

def main():
    n = int(sys.stdin.readline().strip())
    res = []
    available_nums = [x for x in range(1, n+1)]

    for num in available_nums[::-1]:
        variant = []
        max = n
        i = num
        print(num)
        while sum(variant) != n:
            if (sum(variant) + i) <= n:
                variant.append(i)
            elif (sum(variant) + i-1):
                i -= 1
            else:
                break
        res.append(variant)
    print(res)

if __name__ == "__main__":
    main()
