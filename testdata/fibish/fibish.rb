def f(x)
  return 0, x if x < 3
  a, b = f(x - 2)
  c, d = f(x - 1)
  return a + d, b + c
end

x = ARGV.empty? ? 40 : Integer(ARGV[0])
a, b = f(x)
puts "f(#{x})=#{a},#{b}"
