package chapter1

func Fibonazi(n int) int {
	p, q := 0, 1

	if n <= 1 {
		return 0
	}
	if n == 2 {
		return q
	}

	for i := 0; i < n-2; i++ {
		p, q = q, p+q
	}

	return q
}
