# frozen_string_literal: true

# ==============================================
# Array Search in Ruby
#
# This file demonstrates:
# - Linear search (sequential search)
# - Binary search (on sorted arrays)
# - Benchmarking the performance difference
#
# The goal is to understand:
# - How to implement search algorithms by hand
# - Their time complexity (O(N) vs O(log N))
# - How to measure their performance in Ruby
# ==============================================

require 'benchmark'

# ----------------------------------------------
# Linear Search
# ----------------------------------------------

# Linear Search (Index version)
#
# Returns the index of the first element that matches the target.
# Returns nil if the target is not found.
#
# Time Complexity:
# - Best case:    O(1)   (target is at the first position)
# - Average case: O(N)
# - Worst case:   O(N)   (target is at the end or not present)
def linear_search_index(array, target)
  array.each_with_index do |value, index|
    return index if value == target
  end
  nil
end

# Linear Search (Value version)
#
# Returns the value itself if found, otherwise nil.
# This behaves similarly to Ruby's Array#find.
def linear_search_value(array, target)
  array.each do |value|
    return value if value == target
  end
  nil
end

# ----------------------------------------------
# Binary Search
# ----------------------------------------------

# Binary Search (Index version)
#
# Preconditions:
# - The array must be sorted in ascending order.
#
# Returns the index of the target if found, otherwise nil.
#
# Time Complexity:
# - Best case:    O(1)
# - Average case: O(log N)
# - Worst case:   O(log N)
def binary_search_index(array, target)
  left = 0
  right = array.length - 1

  while left <= right
    mid = (left + right) / 2
    value = array[mid]

    if value == target
      return mid
    elsif value < target
      left = mid + 1
    else
      right = mid - 1
    end
  end

  nil
end

# Binary Search (using Ruby's built-in bsearch_index)
#
# This is just for comparison with our manual implementation.
# It requires Ruby 2.3+.
def ruby_bsearch_index(array, target)
  array.bsearch_index { |value| value <=> target }
end

# ----------------------------------------------
# Demo: correctness check
# ----------------------------------------------
def demo_correctness
  puts '== Demo: Correctness Check =='
  array = [1, 3, 5, 7, 9, 11]

  puts '-- Linear Search --'
  p linear_search_index(array, 7)   # => 3
  p linear_search_index(array, 10)  # => nil

  puts '-- Binary Search (manual) --'
  p binary_search_index(array, 7)   # => 3
  p binary_search_index(array, 10)  # => nil

  puts '-- Binary Search (Ruby bsearch_index) --'
  p ruby_bsearch_index(array, 7)    # => 3
  p ruby_bsearch_index(array, 10)   # => nil
end

# ----------------------------------------------
# Benchmark: Linear vs Binary Search
# ----------------------------------------------
def benchmark_searches
  puts
  puts '== Benchmark: Linear Search vs Binary Search =='

  sizes = [10_000, 100_000, 1_000_000]

  sizes.each do |n|
    puts
    puts "Array size: #{n}"
    # Create a sorted array: [0, 1, 2, ..., n-1]
    array = (0...n).to_a

    # Choose a target that is at the very end to simulate
    # the worst case for linear search.
    target = n - 1

    # Run each search multiple times to get a more stable measurement.
    iterations = 50

    Benchmark.bm(20) do |x|
      x.report('linear_search_index') do
        iterations.times { linear_search_index(array, target) }
      end

      x.report('binary_search_index') do
        iterations.times { binary_search_index(array, target) }
      end

      x.report('ruby bsearch_index') do
        iterations.times { ruby_bsearch_index(array, target) }
      end
    end
  end
end

# ----------------------------------------------
# Main
# ----------------------------------------------
if $PROGRAM_NAME == __FILE__
  demo_correctness
  benchmark_searches
end
