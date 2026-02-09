# frozen_string_literal: true

# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val = 0, left = nil, right = nil)
#         @val = val
#         @left = left
#         @right = right
#     end
# end
# @param {Integer[]} nums
# @return {TreeNode}
def sorted_array_to_bst(nums)
  build(nums, 0, nums.length - 1)
end

def build(nums, left, right)
  return nil if left > right

  mid = left + (right - left) / 2
  node = TreeNode.new(nums[mid])

  node.left = build(nums, left, mid - 1)
  node.right = build(nums, mid + 1, right)
  node
end
