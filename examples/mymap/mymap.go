package mymap

func count(s string, m map[rune]int) {
	// s := "aaabbBcccDdd"
	// m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	return
}

// HasDupeRune is ...
func HasDupeRune(s string) bool {
	// m1 := make(map[rune]bool)
	// for _, r := range s {
	// 	if m1[r] {
	// 		return true
	// 	}
	// 	m1[r] = true
	// }
	m2 := make(map[rune]struct{})
	for _, r := range s {
		if _, ok := m2[r]; ok {
			return true
		}
		m2[r] = struct{}{}
	}
	return false
}
