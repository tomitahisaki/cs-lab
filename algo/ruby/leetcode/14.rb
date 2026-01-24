# # frozen_string_literal: true

# @params {String[]} strs
# @returns {String}

strs = ["flower","flow","flight"]

def longest_common_prefix(strs)
  return "" if strs.empty?

  prefix = strs[0]
  strs[1..].each do |str|
    while !str.start_with?(prefix)
      prefix = prefix[0..-2]
      return "" if prefix.empty?
    end
  end
  prefix
end

puts longest_common_prefix(strs)
