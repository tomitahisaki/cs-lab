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
# @return {Boolean}
def is_balanced(root)
  height(root) != -1
end

def height(node)
  return 0 if node.nil?

  left_height = height(node.left)
  return -1 if left_height == -1

  right_height = height(node.right)
  return -1 if right_height == -1

  return -1 if (left_height - right_height).abs > 1

  1 + [left_height, right_height].max
end
