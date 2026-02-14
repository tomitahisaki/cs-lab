# frozen_string_literal: true
triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
# @param {Integer[][]} triangle
# @return {Integer}
def minimum_total(triangle)
  dp = triangle[-1].dup

  (triangle.length - 2).downto(0) do |i|
    (0..i).each do |j|
      dp[j] = triangle[i][j] + [dp[j], dp[j + 1]].min
    end
  end
  dp[0]
end

p minimum_total(triangle) # 11
