local function f(x, unused)
    if x < 3 then
        return 0, x
    end
    local a, b = f(x - 2, unused)
    local c, d = f(x - 1, unused)
    return a + d, b + c
end

local x = tonumber(arg[1]) or 40
local a, b = f(x, f)
print(string.format("f(%d)=%d,%d", x, a, b))
