package main

import "fmt"

func main() {
	var n, i uint

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&i)
		fmt.Println(i ^ 0xFFFF)

	}

}
