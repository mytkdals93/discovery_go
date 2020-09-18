package main

import "fmt"

func main() {
	a := "안녕하세요"
	n := "Hello"
	s := rune(44032) // 가
	e := rune(44060) // 갛 다음
	//유니코드에 대하여 바이트 위치는 i, 각 문자는 r에 넣음
	// i: int , r: rune 타입
	//rune타입은 int32 (4byte의 정수의 별칭) char타입은 1byte
	// i의 크기가 0, 3, 6... 이 되는 이유는 UTF-8로 한글을 표현하는데 3byte가 필요함
	for i, r := range a {
		fmt.Println(i, r, string(r))
	}

	for i, r := range n {
		fmt.Println(i, r)
	}

	for i := s; i <= e; i++ {
		fmt.Println(i, string(i))
	}
}
