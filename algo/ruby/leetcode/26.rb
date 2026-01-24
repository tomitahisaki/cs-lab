# frozen_string_literal: true

# @params {Integer[]} nums
# @returns {Integer}

def remove_duplicates(nums)
  nums.uniq! # 重複を削除がないとnilになるので、メソッドチェーンしない
  nums.length
end
