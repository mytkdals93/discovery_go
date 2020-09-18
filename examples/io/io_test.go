package io

import (
	"fmt"
	"os"
	"strings"
)

func Example_WriteTo() {
	lines := []string{
		"bill@mail.com",
		"tom@mail.com",
		"jane@mail.com",
		"sangmin@mail.com",
	}
	if err := WriteTo(os.Stdout, lines); err != nil {
		fmt.Println(err)
	}
	//Output:
	//bill@mail.com
	//tom@mail.com
	//jane@mail.com
	//sangmin@mail.com
}

func Example_ReadFrom() {
	r := strings.NewReader("bill\ntom\njane\n")
	alines := &[]string{}
	if err := ReadFrom(r, alines); err != nil {
		fmt.Println(err)
	}
	fmt.Println(*alines)
	//Output:
	//[bill tom jane]
}
