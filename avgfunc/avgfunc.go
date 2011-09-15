package main

func avg(nums []float64) (avg float64) {
	sum := 0.0
	switch len(nums) {
	case 0:
		avg = 0
	default:
		for _, n := range nums {
			sum += n
		}

		avg = sum/float64(len(nums))
	}
	return
}

func main() {
	vals := []float64{1.0,2.0,3.0,4.0,5.0}

	println(avg(vals))
}