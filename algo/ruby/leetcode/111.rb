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
# @param {TreeNode} root
# @return {Integer}
def min_depth(root)
  return 0 if root.nil?
  return 1 if root.left.nil? && root.right.nil?

  left_depth = min_depth(root.left)
  right_depth = min_depth(root.right)

  if root.left.nil? || root.right.nil?
    1 + left_depth + right_depth
  else
    1 + [left_depth, right_depth].min
  end
end

