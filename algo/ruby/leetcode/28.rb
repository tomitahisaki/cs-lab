# frozen_string_literal: true

# @param {String} haystack
# @param {String} needle
# @return {Integer}
def str_str(haystack, needle)
  haystack.index(needle) || -1
end

puts str_str(haystack, needle) # => 0
