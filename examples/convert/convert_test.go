package convert

import (
	"fmt"
	"strconv"
)

func Example_Convert() {
	var (
		i   int
		t   int64
		k   int64
		f   float64
		si  string
		s0x string
	)
	//string에서 숫자로 변환
	i, _ = strconv.Atoi("350")                 //i = 350
	t, _ = strconv.ParseInt("cc7fdd", 16, 64)  // t = 13402077
	k, _ = strconv.ParseInt("0xcc7fdd", 0, 64) // k = 13402077
	f, _ = strconv.ParseFloat("3.14", 64)      // f = 13402077
	//숫자에서 string으로 변환
	si = strconv.Itoa(340)                // s = "340"
	s0x = strconv.FormatInt(13402077, 16) // s = "cc7fdd"

	fmt.Printf("%#v \n", i)
	fmt.Printf("%#b \n", t)
	fmt.Printf("%#b \n", k)
	fmt.Printf("%#v \n", f)
	fmt.Printf("%#v \n", si)
	fmt.Printf("%#v \n", s0x)
	//Output:
	//
}
