n, q = gets.chomp.split.map(&:to_i)
follows_map = Array.new(n) { Array.new(n, 'N') }

def follow(follows_map, a, b)
  follows_map[a - 1][b - 1] = 'Y'
end

def follow_all(follows_map, a)
  n = follows_map.length
  0.upto(n - 1) do |b|
    follows_map[a - 1][b] = 'Y' if follows_map[b][a - 1] == 'Y'
  end
end

def follow_follow(follows_map, a)
  n = follows_map.length
  a_idx = a - 1

  # クエリ開始時点の a のフォロー状況を固定（ここが重要）
  original = follows_map[a_idx].dup

  0.upto(n - 1) do |b|
    next unless original[b] == 'Y' # ← 途中で増えた分は見ない

    0.upto(n - 1) do |c|
      next if c == a_idx

      follows_map[a_idx][c] = 'Y' if follows_map[b][c] == 'Y'
    end
  end
end

1.upto(q) do |i|
  array = gets.chomp.split.map(&:to_i)
  if array[0] == 1
    follow(follows_map, array[1], array[2])
    puts(follows_map.map { |row| row.join })
    puts 'follow==========================='
  elsif array[0] == 2
    follow_all(follows_map, array[1])
    puts(follows_map.map { |row| row.join })
    puts 'follow_all==========================='
  else
    follow_follow(follows_map, array[1])
    puts(follows_map.map { |row| row.join })
    puts 'follow_follow==========================='
  end
end

puts(follows_map.map { |row| row.join })
