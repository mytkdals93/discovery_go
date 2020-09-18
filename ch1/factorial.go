package c1

func facItr1(n int) int {
	result := 1
	for n > 0 {
		result *= n
		n--
	}
	return result
}

func facItr2(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
