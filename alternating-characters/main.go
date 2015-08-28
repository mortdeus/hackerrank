package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {

		var s string
		var ndel = 0

		fmt.Scan(&s)

		for i, c := range s {
			if i+1 == len(s) {
				continue
			}
			if c == rune(s[i+1]) {
				ndel++
			}
		}
		fmt.Println(ndel)
	}
}
