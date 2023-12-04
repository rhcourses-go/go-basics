package counting

func CountSomething(n int) int {
	result := 0
	for i := 0; i < n; i++ {
		if i%3 == 0 {
			result++
		}
	}
	return result
}
