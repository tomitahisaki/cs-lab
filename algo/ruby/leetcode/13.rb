# # frozen_string_literal: true

s = "III"

# @params {String} s
# @returns {Integer}
def roman_to_int(string)
  symbols = {
    I: 1,
    V: 5,
    X: 10,
    L: 50,
    C: 100,
    D: 500,
    M: 1000
  }
  result = 0

  string.each_char.with_index do |char, index|
    value = symbols[char.to_sym]

    if index < string.length - 1 &&
       value < symbols[string[index + 1].to_sym]
      result -= value
    else
      result += value
    end
  end

  result
end

puts roman_to_int(s)
