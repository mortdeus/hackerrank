package main

import "fmt"

var maxtab = []maxab{{max: 0, pair: [2]int{0, 0}}}

type maxab struct {
	max  int
	pair [2]int
}

func main() {
	var l, r, xab int

	fmt.Scan(&l, &r)

	for a := l; a <= r; a++ {
		for b := a; b <= r; b++ {

			xab = a ^ b
			if m := maxtab[len(maxtab)-1].max; m < xab {
				maxtab = append(maxtab, maxab{
					max:  xab,
					pair: [2]int{a, b},
				})
			}
		}
	}

	fmt.Println(maxtab[len(maxtab)-1].max)
}
