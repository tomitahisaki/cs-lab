n = gets.chomp.to_i
lines = n.times.map { gets.chomp.to_i }

expected_sum = n * (n + 1) / 2
actual_sum = lines.sum

if expected_sum == actual_sum
  puts 'Correct'
else
  dupulicated_number = lines.tally.find { |_, v| v > 1 }[0]
  missing_number = dupulicated_number + (expected_sum - actual_sum)
  puts "#{dupulicated_number} #{missing_number.abs}"
end
