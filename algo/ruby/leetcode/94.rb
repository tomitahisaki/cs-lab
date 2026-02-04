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
# @return {Integer[]}
def inorder_traversal(root)
  result = []  
  inorder(root, result)
  result
end

def inorder(node, result)
  return if node.nil?

  inorder(node.left, result) # 左の子ノードをnilになるまで辿る
  result << node.val
  inorder(node.right, result) # 右の子ノードをnilになるまで辿る
end

p inorder_traversal(root)
