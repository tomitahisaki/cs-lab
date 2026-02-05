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
def is_symmetric(root)
  left_tree = root.left
  right_tree = root.right

  is_mirror(left_tree, right_tree)
end

def is_mirror(left_tree, right_tree)
  return true if left_tree.nil? && right_tree.nil?
  return false if left_tree.nil? || right_tree.nil?
  return false if left_tree.val != right_tree.val

  is_mirror(left_tree.left, right_tree.right) && is_mirror(left_tree.right, right_tree.left)
end
