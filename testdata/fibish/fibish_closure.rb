def f(x, unused)
  return 0, x if x < 3
  a, b = f(x - 2, unused)
  c, d = f(x - 1, unused)
  return a + d, b + c
end

x = ARGV.empty? ? 40 : Integer(ARGV[0])
a, b = f(x, method(:f))
puts "f(#{x})=#{a},#{b}"
