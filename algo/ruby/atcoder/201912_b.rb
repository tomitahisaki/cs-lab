n = gets.chomp.to_i
lines = n.times.map { gets.chomp.to_i }

(lines.count - 1).times do |i|
  if lines[i] < lines[i + 1]
    puts "up #{lines[i + 1] - lines[i]}"
  elsif lines[i] > lines[i + 1]
    puts "down #{lines[i] - lines[i + 1]}"
  else
    puts 'stay'
  end
end
