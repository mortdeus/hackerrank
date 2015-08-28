package main

import "fmt"

func main() {
	var t, n int

	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		fmt.Println("growth(n):", growth(cycle(n)))
		fmt.Println("growthRecursion(n):", growthRecursion(cycle(n)))
	}
}

type cycle int
type height int

var growthTable = map[cycle]height{0: 1}

func growthRecursion(c cycle) (h height) {
	var exists bool
	if h, exists = growthTable[c]; exists {
		return
	}
	if c%2 == 0 {
		h = growth(c-1) + 1

		growthTable[c] = h
		return
	}
	if c%2 == 1 {
		h = growth(c-1) * 2
		growthTable[c] = h
		return
	}
	panic("invalid c")
}

func growth(c cycle) height {
	return -(-1 << (c >> 1)) << c % 2
}
