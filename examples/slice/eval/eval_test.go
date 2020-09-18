package eval

import "fmt"

func ExampleEval() {
	fmt.Println(Eval("5"))
	fmt.Println(Eval("5 + 2"))
	fmt.Println(Eval("5 + 2 * 3"))
	fmt.Println(Eval("5 * ( 3 + 1 ) / 2"))

	//Output:
	//5
	//7
	//11
	//10
}
