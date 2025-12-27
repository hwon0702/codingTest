from math import gcd
tc = int(input())
for i in range(tc):
    a,b = map(int, input().split())
    getGcd = gcd(a,b)
    print('{:.0f}'.format(a*b/getGcd))