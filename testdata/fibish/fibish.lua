local function f(x)
    if x < 3 then
        return 0, x
    end
    local a, b = f(x - 2)
    local c, d = f(x - 1)
    return a + d, b + c
end

local x = tonumber(arg[1]) or 40
local a, b = f(x)
print(string.format("f(%d)=%d,%d", x, a, b))
