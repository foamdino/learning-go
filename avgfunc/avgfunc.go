package main

func avg(nums []float64) (avg float64) {
	sum := 0.0
	for _, n := range nums {
		sum += n
	}

	return sum/float64(len(nums))
}

func main() {
	vals := []float64{1.0,2.0,3.0,4.0,5.0}

	println(avg(vals))
}