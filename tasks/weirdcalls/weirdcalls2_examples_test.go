package weirdcalls

import "fmt"

func ExampleFoo1() {
	Foo1()

	// Output:
	// 19
	// 9
	// 5
}

func ExampleFoo2() {
	Foo2(5)

	// Output:
	// 19
	// 9
}

func ExampleFoo3() {
	fmt.Println(Foo3(15))

	// Output:
	// 19
	// 9
}
