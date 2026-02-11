# frozen_string_literal: true

# @param {Integer} num_rows
# @return {Integer[][]}
def generate(num_rows)
  result = []
  return result if num_rows == 0

  result << [1]

  (1...num_rows).each do |i|
    prev_row = result[i - 1]
    row = [1]

    (1...prev_row.length).each do |j|
      row << prev_row[j - 1] + prev_row[j]
    end

    row << 1
    result << row
  end

  result
end

p generate(5)

