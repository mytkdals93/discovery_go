package c1

func Fibonazi(n int) int {
	p, q := 0, 1

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	for i := 0; i < n-1; i++ {
		p, q = q, p+q
	}

	return q
}
