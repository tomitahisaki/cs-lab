# frozen_string_literal: true

# Linear Search in Ruby (index version)
#
# Returns the index of the first element that matches the target.
# Returns nil if the target is not found.
# This is equivalent to Ruby's Array#find_index.
#
# Time Complexity: O(N)
def linear_search(array, target)
  array.each_with_index do |value, index|
    return index if value == target
  end
  nil
end

# --- Demo ---
# puts linear_search([10, 20, 30, 40, 50], 30) # Output: 2

# Linear Search in Ruby (value version)
#
# Returns the *value* of the first element that matches the target.
# Returns nil if the target is not found.
# This is equivalent to Ruby's Array#find.
#
# Time Complexity: O(N)
def linear_search_value(array, target)
  array.each do |value|
    return value if value == target
  end
  nil
end

# --- Demo ---
puts linear_search_value([10, 20, 30, 40, 50], 30) # Output: 30
