# ==============================================
# Ruby Reference Study
# cs/ruby/reference/reference_study.rb
#
# This file demonstrates how Ruby handles:
# - variables and object references
# - destructive vs non-destructive operations
# - method argument passing (pass-by-value of references)
# - differences between reassignment and mutation
# ==============================================

puts '== 1. Variables and Object IDs =='
a = 'hello'
b = a

puts "a.object_id: #{a.object_id}"
puts "b.object_id: #{b.object_id}"
puts '→ Both variables reference the same object.'
puts

puts '== 2. Destructive method modifies the object itself =='
str = 'ruby'
puts "before: #{str}, object_id: #{str.object_id}"

# upcase! modifies the original string (destructive)
str.upcase!
puts "after:  #{str}, object_id: #{str.object_id}"
puts '→ Destructive methods change the underlying object.'
puts

puts '== 3. Mutation affects all references =='
x = [1, 2, 3]
y = x
puts "before x: #{x}, y: #{y}"

# << mutates the array (does not reassign)
y << 4
puts "after  x: #{x}, y: #{y}"
puts "object_id x: #{x.object_id}, y: #{y.object_id}"
puts '→ Both variables point to the same array, so mutation is shared.'
puts

puts '== 4. Method receives a copied reference =='
def modify(arr)
  # This mutates the original array object
  arr << 999
end

list = [1, 2, 3]
puts "before: #{list}, object_id: #{list.object_id}"
modify(list)
puts "after:  #{list}, object_id: #{list.object_id}"
puts '→ Ruby passes references by value. The object can still be mutated.'
puts

puts '== 5. Reassignment inside a method does NOT affect caller =='
def reassign(a)
  # Reassigns the local variable 'a' to a new object
  a = 'changed'
  puts "  inside method: #{a} (object_id: #{a.object_id})"
end

text = 'original'
puts "before: #{text} (object_id: #{text.object_id})"
reassign(text)
puts "after:  #{text} (object_id: #{text.object_id})"
puts '→ Reassignment only affects the local variable, not the caller.'
puts

puts '== 6. Hash mutation behaves the same way as arrays =='
h1 = { a: 1 }
h2 = h1
puts "before: #{h1}"

# Modifying the hash affects all references
h2[:b] = 2
puts "after h1: #{h1}, h2: #{h2}"
puts "object_id h1: #{h1.object_id}, h2: #{h2.object_id}"
puts '→ Both variables reference the same hash object.'
puts

puts '== 7. Array object_id vs element object_id =='
arr = [1, 2, 3]
puts "arr.object_id: #{arr.object_id}"
puts "arr[0].object_id: #{arr[0].object_id}"
puts '→ Array object and its elements are separate objects.'
