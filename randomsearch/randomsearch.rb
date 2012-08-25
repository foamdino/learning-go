def objective_function(vector)
	return vector.inject(0) {|sum, x| sum + (x ** 2.0)}
end

def random_vector(minmax)
	return Array.new(minmax.size) do |i|
		minmax[i][0] + ((minmax[i][1] - minmax[i][0]) * rand())
	end
end

def search(search_space, max_iter)
	best = nil

	max_iter.times do |iter|
		candidate = {}
		candidate[:vector] = random_vector(search_space)
		candidate[:cost] = objective_function(candidate[:vector])
		best = candidate if best.nil? or candidate[:cost] < best[:cost]
		puts " > iteration=#{(iter+1)}, best=#{best[:cost]}, #{best[:vector].inspect}"
	end

	return best
end

if __FILE__ == $0
	problem_size = 2
	search_space = Array.new(problem_size) {|i| [-5, 5]}
	max_iter = 100
	best = search(search_space, max_iter)
	puts "Done. Best solution: c=#{best[:cost]}, v=#{best[:vector].inspect}"
end