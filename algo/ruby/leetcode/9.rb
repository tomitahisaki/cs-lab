# frozen_string_literal: true

# @params {Integer} x
# @returns {Boolean}
def is-palindrome(x)
  return false if x.negative?

  str = x.to_s
  str == str.reverse
end
