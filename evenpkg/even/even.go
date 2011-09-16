package even

// functions with Uppercase are exported
func Even(i int) bool {
	return i % 2 == 0
}

// function in lowercase are private
func odd(i int) bool {
	return i % 2 == 1
}