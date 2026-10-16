import sys


# A list's first slot represents *int; sharing lists preserves pointer aliases.
def f(x, xm1, xm2, p):
    y = [x - 4, 0]
    if x < 2:
        p[0] += x
        return
    xp = [x - 3]
    g(xm1[0], xm2, xp, p)
    h(xm2[0], xp, y, p)


def g(x, xm1, xm2, p):
    y = [x - 4, 0, 0]
    if x < 2:
        p[0] += x
        return
    xp = [x - 3]
    k(xm2[0], xp, y, p)
    h(xm1[0], xm2, xp, p)


def h(x, xm1, xm2, p):
    y = [x - 4, 0, 0, 0]
    if x < 2:
        p[0] += x
        return
    xp = [x - 3]
    k(xm1[0], xm2, xp, p)
    f(xm2[0], xp, y, p)


def k(x, xm1, xm2, p):
    y = [x - 4, 0, 0, 0, 0]
    if x < 2:
        p[0] += x
        return
    xp = [x - 3]
    f(xm2[0], xp, y, p)
    g(xm1[0], xm2, xp, p)


x = int(sys.argv[1]) if len(sys.argv) > 1 else 40
y, xm1, xm2 = [0], [x - 1], [x - 2]
f(x, xm1, xm2, y)
print(f"Fib({x})={y[0]}")
