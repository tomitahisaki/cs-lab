# frozen_string_literal: true

# Definition for singly-linked list.
class ListNode
    attr_accessor :val, :next
    def initialize(val = 0, _next = nil)
        @val = val
        @next = _next
    end
end
# @param {ListNode} list1
# @param {ListNode} list2
# @return {ListNode}
def merge_two_lists(list1, list2)
  result = ListNode.new(0)
  current = result

  while list1 && list2
    if list1.val <= list2.val
      current.next = list1
      list1 = list1.next
    else
      current .next = list2
      list2 = list2.next
    end
    current = current.next
  end

  current.next = list1 || list2

  return result.next
end

p merge_two_lists(list1, list2)



