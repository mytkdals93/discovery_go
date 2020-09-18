package slice

import "fmt"

func Example_slicing() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])

	//Output:
	//[1 2 3 4 5]
	//[2 3]
	//[3 4 5]
	//[1 2 3]
}

func Example_append() {
	f1 := []string{"사과", "바나나", "토마토"}
	f2 := []string{"포도", "수박", "아보카도"}
	f3 := append(f1, f2...)
	f4 := append(f1[:2], f2...)
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)
	//Output:
	//[사과 바나나 토마토]
	//[포도 수박 아보카도]
	//[사과 바나나 토마토 포도 수박 아보카도]
	//[사과 바나나 포도 수박 아보카도]

}

func Example_sliceCap() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println("len: ", len(nums))
	fmt.Println("cap: ", cap(nums))
	fmt.Println()

	sliced1 := nums[:3]
	fmt.Println(sliced1)
	fmt.Println("len: ", len(sliced1))
	fmt.Println("cap: ", cap(sliced1))
	fmt.Println()

	sliced2 := nums[:2]
	fmt.Println(sliced2)
	fmt.Println("len: ", len(sliced2))
	fmt.Println("cap: ", cap(sliced2))
	fmt.Println()

	sliced3 := sliced1[:4]
	fmt.Println(sliced3)
	fmt.Println("len: ", len(sliced3))
	fmt.Println("cap: ", cap(sliced3))
	fmt.Println()

	nums[2] = 100
	fmt.Println(nums, sliced1, sliced2, sliced3)
	//Output:
	//[1 2 3 4 5]
	//len:  5
	//cap:  5
	//
	//[1 2 3]
	//len:  3
	//cap:  5
	//
	//[1 2]
	//len:  2
	//cap:  5
	//
	//[1 2 3 4]
	//len:  4
	//cap:  5
	//
	//[1 2 100 4 5] [1 2 100] [1 2] [1 2 100 4]
}
func Example_sliceCopy() {
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, len(src))
	// for i := range src {
	// 	dest[i] = src[i]
	// }
	// fmt.Println(dest)
	// //Output: [30 20 50 10 40]

	//복사할 경우 사용하는 copy함수 둘중에 길이가 더 적은것에 맞춰서 반환
	//복사한 길이만큼 반환
	// if n := copy(dest, src); n != len(src) {
	// 	fmt.Println("복사가 덜 됐습니다.")
	// }
	// fmt.Println(dest)
	// //Output: [30 20 50 10 40]
	dest = append([]int{}, src...)
	fmt.Println(dest)
	//Output: [30 20 50 10 40]

}
func Example_slicePut() {
	// a의 중간에 x정수를 끼워넣고 싶은경우
	// ex) a[2]에 100을 끼워넣고 싶다 => [0 1 100 2 3 4 5]
	i := 2
	x := 100
	a := []int{0, 1, 2, 3, 4, 5}
	//i가 len(a)와 크거나 같다면 i+1을 해주기 떄문에 인덱스를 초과하여 에러발생
	if i < len(a) {
		a = append(a[:i+1], a[i:]...)
	} else {
		a = append(a, x)
	}
	a[i] = x
	fmt.Println(a)
	//Output: [0 1 100 2 3 4 5]
}
func Example_slicePut2() {
	// a의 중간에 x정수를 끼워넣고 싶은경우
	// ex) a[2]에 100을 끼워넣고 싶다 => [0 1 100 2 3 4 5]
	a := []int{0, 1, 2, 3, 4, 5}
	i := len(a)
	x := 100

	a = append(a, x)
	copy(a[i+1:], a[i:])
	a[i] = x
	fmt.Println(a)
	//배열을 추가하고 싶은겨우에도 사용가능
	i = 3
	xs := []int{7, 8, 9}
	a = append(a, xs...)
	copy(a[i+len(xs):], a[i:])
	copy(a[i:], xs)
	fmt.Println(a)
	//Output:
	//[0 1 2 3 4 5 100]
	//[0 1 2 7 8 9 3 4 5 100]
}

func Example_sliceDelete() {
	// a의 중간에 i에 있는 엘리멘트를 삭제하고싶다
	// ex) a[2]를 삭제하고 싶은경우 => [0 1 3 4 5 ]
	a := []int{0, 1, 2, 3, 4, 5}
	i := 2

	a = append(a[:i], a[i+1:]...)
	fmt.Println(a)

	// index에서 k개를 지우고 싶다.
	a = []int{0, 1, 2, 3, 4, 5}
	k := 2
	a = append(a[:i], a[i+k:]...)
	fmt.Println(a)

	//순서가 상관없이 해당 인덱스를 지우고 싶은경우
	a = []int{0, 1, 2, 3, 4, 5}
	i = 1
	a[i] = a[len(a)-1]
	a = a[:len(a)-1]
	fmt.Println(a)

	//순서가 상관없이 해당 인덱스부터 k개를 지우고 싶은경우
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	i = 3
	k = 4
	copy(a[i:], a[i+k:])
	// 정수형이 아니라 포인터형 슬라이스라면 메모리 누수를 방지해야함.
	// for i := 0; i < k; i++ {
	// 	a[len(a)-1-i] = nil
	// }
	a = a[:len(a)-k]
	fmt.Println(a)
	//Output:
	//[0 1 3 4 5]
	//[0 1 4 5]
	//[0 5 2 3 4]
	//[0 1 2 7 8 9]
}
