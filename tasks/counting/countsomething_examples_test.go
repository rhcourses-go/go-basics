package counting

import "fmt"

func ExampleCountSomething() {
	// Sinnvolle Aufrufe von CountSomething()
	fmt.Println(CountSomething(15))
	fmt.Println(CountSomething(25))
	fmt.Println(CountSomething(35))

	// Unsinnige Aufrufe von CountSomething()
	fmt.Println(CountSomething(0))
	fmt.Println(CountSomething(-1))
	fmt.Println(CountSomething(-2))

	// Output:
	// 5
	// 9
	// 12
	// 0
	// 0
	// 0
}
