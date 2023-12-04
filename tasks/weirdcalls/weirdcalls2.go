package weirdcalls

import "fmt"

func Foo1() {
	x := 5
	Foo2(x)
	fmt.Println(x)
}

func Foo2(x int) {
	x = x * 3
	y := Foo3(x)
	fmt.Println(y)
}

func Foo3(x int) int {
	x = x + 4
	fmt.Println(x)
	return x / 2
}
