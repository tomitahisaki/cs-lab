# frozen_string_literal: true

# @param {Integer} row_index
# @return {Integer[]}
def get_row(row_index)
  result = []
  result << [1]

  # row_index is 0-based, so we need to generate rows up to row_index
  (1...row_index+1).each do |i|
    prev_row = result[i - 1]
    row = [1]

    (1...prev_row.length).each do |j|
      row << prev_row[j - 1] + prev_row[j]
    end

    row << 1
    result << row
  end

  result[row_index]
end

p get_row(3) # [1, 3, 3, 1]
