s = gets.chomp
words = []

word = ''
count = 0
s.each_char do |value|
  count += 1 if value.match?(/[A-Z]/)
  word += value

  next unless count == 2

  words << word
  count = 0
  word = ''
end

puts words.sort_by(&:downcase).join('')
