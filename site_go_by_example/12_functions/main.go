package main

import "fmt"

func main() {
	result := a(8)
	fmt.Println(result)
	fmt.Printf("variadic: %v\n", variadic(4, 5, 6))
	runCounters()
	runRecursion(3)
}

func a(x int) string {
	i, _, s := b(x, 0)
	fmt.Printf("in a: %d\n", i)
	return fmt.Sprintf("a: (%s)", s)
}

func b(x, y int) (int, int, string) {
	return 1, 7, fmt.Sprintf("b: (%d and %d)", x, y)
}

func variadic(ints ...int) map[int]int {
	result := make(map[int]int)
	var sum int
	for i, value := range ints {
		result[i] = value
		sum += value
	}
	result[len(result)] = sum
	return result
}

func runCounters() {
	counter1 := newCounter()
	counter2 := newCounter()
	counter1()
	counter1()
	fmt.Println(counter1())
	fmt.Println(counter2())
}

func newCounter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func runRecursion(level int) {

	if level <= 0 {
		fmt.Println("Recursion level 0.")
		return
	}

	fmt.Printf("Recursion level %d -> ", level)
	runRecursion(level - 1)
}
