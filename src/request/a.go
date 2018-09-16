package request

import "fmt"

func A() {
	var a = 2
	if a > 1 {
		a := 5
		fmt.Println(a)
	} else {
		var a = 7
		fmt.Println(a)
	}
	if a > 1 {
		var a = 7
		fmt.Println(a)
	}
	fmt.Println(a);
}