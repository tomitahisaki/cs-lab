# # frozen_string_literal: true

# @params {String} s
# @returns {Boolean}

s = '([)]'

def is_valid(s)
  stack = []
  pair = {
    '(' => ')',
    '[' => ']',
    '{' => '}'
  }
  s.each_char do |char|
    if pair.key?(char)
      stack.push(char)
    else
      return false if stack.empty?

      last_open = stack.pop
      return false if pair[last_open] != char
    end
  end

  stack.empty?
end

puts is_valid(s)
