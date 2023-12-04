package weirdcalls

import "fmt"

func ExampleBar() {
	fmt.Println(Bar(4)) // 4*3+4 = 16; 16/2 = 8
	fmt.Println(Bar(5)) // 5*3+4 = 19; 19/2 = 9
	fmt.Println(Bar(6)) // 6*3+4 = 22; 22/2 = 11

	// Output:
	// 8
	// 9
	// 11
}

func ExampleFoo() {
	Foo()

	// Output:
	// 15
}
