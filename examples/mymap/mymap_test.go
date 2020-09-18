package mymap

import (
	"fmt"
	"sort"
)

func Example_count() {
	s := "aaabbBcccDdd"
	m := make(map[rune]int)

	count(s, m)
	var keys sort.IntSlice

	for key := range m {
		keys = append(keys, int(key))
	}
	sort.Sort(keys)
	for _, key := range keys {
		fmt.Println(string(rune(key)), m[rune(key)])
	}
	//Output:
	//B 1
	//D 1
	//a 3
	//b 2
	//c 3
	//d 2
}
func Example_HasDupeRune() {
	s1 := "안녕하세요"
	s2 := "다시합시다"
	fmt.Println(HasDupeRune(s1))
	fmt.Println(HasDupeRune(s2))
	//Output:
	//false
	//true
}
func Example_DeleteMap() {
	temp := make(map[string]int)
	temp["a"] = 0
	temp["b"] = 1

	fmt.Println(temp)
	delete(temp, "a")
	fmt.Println(temp)
	//Output:
	//map[a:0 b:1]
	//map[b:1]
}
