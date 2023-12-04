package counting

import "fmt"

func ExampleCountElements() {
	l1 := []int{3, 4}
	l2 := []int{3}
	l3 := []int{5}
	l4 := []int{3, 4, 5}

	fmt.Println(CountElements(l1))
	fmt.Println(CountElements(l2))
	fmt.Println(CountElements(l3))
	fmt.Println(CountElements(l4))

	// Output:
	// 25
	// 9
	// 25
	// 50
}
