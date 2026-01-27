# frozen_string_literal: true

s = "   fly me   to   the moon  "
# @param {String} s
# @return {Integer}
def length_of_last_word(s)
  result = s.strip.split(" ")

  result[-1].length
end

puts length_of_last_word(s) # => 5
