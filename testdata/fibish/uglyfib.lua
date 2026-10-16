-- Translation of uglyfib.go. A table's first slot represents *int;
-- passing the same table preserves pointer aliasing. Arrays retain their size.
local f, g, h, k

f = function(x, xm1, xm2, p)
    local y = {x - 4, 0}
    if x < 2 then
        p[1] = p[1] + x
        return
    end
    local xp = {x - 3}
    g(xm1[1], xm2, xp, p)
    h(xm2[1], xp, y, p)
end

g = function(x, xm1, xm2, p)
    local y = {x - 4, 0, 0}
    if x < 2 then
        p[1] = p[1] + x
        return
    end
    local xp = {x - 3}
    k(xm2[1], xp, y, p)
    h(xm1[1], xm2, xp, p)
end

h = function(x, xm1, xm2, p)
    local y = {x - 4, 0, 0, 0}
    if x < 2 then
        p[1] = p[1] + x
        return
    end
    local xp = {x - 3}
    k(xm1[1], xm2, xp, p)
    f(xm2[1], xp, y, p)
end

k = function(x, xm1, xm2, p)
    local y = {x - 4, 0, 0, 0, 0}
    if x < 2 then
        p[1] = p[1] + x
        return
    end
    local xp = {x - 3}
    f(xm2[1], xp, y, p)
    g(xm1[1], xm2, xp, p)
end

local x = tonumber(arg[1]) or 40
local y, xm1, xm2 = {0}, {x - 1}, {x - 2}
f(x, xm1, xm2, y)
print(string.format("Fib(%d)=%d", x, y[1]))
