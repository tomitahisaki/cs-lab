# frozen_string_literal: true

a = '1010'
b = '1011'
# @param {String} a
# @param {String} b
# @return {String}
def add_binary(a, b)
  a_int = a.to_i(2)
  b_int = b.to_i(2)
  (a_int + b_int).to_s(2)
end

puts add_binary(a, b)
