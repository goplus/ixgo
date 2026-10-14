import sys


def f(x):
    if x < 3:
        return 0, x
    a, b = f(x - 2)
    c, d = f(x - 1)
    return a + d, b + c


x = int(sys.argv[1]) if len(sys.argv) > 1 else 40
a, b = f(x)
print(f"f({x})={a},{b}")
