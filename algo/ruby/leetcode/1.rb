target = 6
nums = [3, 3]

def two_sum(nums, target)
  hash = {}
  nums.each_with_index do |num, index|
    needed = target - num
    if hash.key?(needed)
      return [hash[needed], index]
    end
    hash[num] = index
  end
end

puts two_sum(nums, target).inspect
