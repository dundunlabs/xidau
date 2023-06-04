package arrays

func Sum(s []int) int {
	sum := 0
	for _, e := range s {
		sum += e
	}
	return sum
}
