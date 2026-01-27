# frozen_string_literal: true

digits = [1, 2, 3]
# @param {Integer[]} digits
# @return {Integer[]}
def plus_one(digits)
  concat_number = digits.join.to_i
  result_number = concat_number + 1
  result_number.to_s.split("").map(&:to_i)
end

p plus_one(digits) # => [1,2,4]

