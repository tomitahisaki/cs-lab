# frozen_string_literal: true

# @param {Integer} n
# @return {Integer}
def climb_stairs(n)
  return 1 if n = 1
  return 2 if n == 2

  prev2 = 1
  prev1 = 2

  (3..n).each do
    current = prev1 + prev2
    prev2 = prev1
    prev1 = current
  end

  prev1
end
