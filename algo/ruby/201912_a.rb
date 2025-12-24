# frozen_string_literal: true

def multiply_by_two(value)
  return 'error' unless value.match?(/\A-?\d+\z/)

  value.to_i * 2
end

n = gets.chomp
result = multiply_by_two(n)

puts result
