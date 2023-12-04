package counting

func CountElements(list []int) int {
	result := 0
	for i := 0; i < len(list); i++ {
		x := list[i]
		result += Value(x)
	}
	return result
}

func Value(x int) int {
	x = x * x
	return x
}
