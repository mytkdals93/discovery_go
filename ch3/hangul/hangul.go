// Package hangul provides basic functions for Hangul processing.
package hangul

var (
	start = rune(44032) // "가"의 유니코드 포인트
	end   = rune(55204) // "힣" 다음 글자의 유니코드 포인트
)

//HasConsonantSuffix returens true if s has Hangul consonant jamo at the end.
func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false
	temp := []rune(s)
	// 마지막 글자
	r := temp[len(temp)-1]
	//마지막 글자가 한글이라면
	if start <= r && r < end {
		index := int(r - start)
		//받침이 있다가 없는 경우가 28번마다 반복됨 가 -(27번)각....-개
		//즉 28번째 마다 받침이 없음 고로 0이면 받침없고 아니면 있는것
		result = index%numEnds != 0
	}

	// for _, r := range s {
	// 	if start <= r && r < end {
	// 		index := int(r - start)
	// 		result = index%numEnds != 0
	// 	}
	// }
	return result
}
