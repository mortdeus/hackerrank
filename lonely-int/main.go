package main

import "fmt"
import "sort"

func main() {
	var narg int

	if _, err := fmt.Scanln(&narg); err != nil {
		panic(err)
	}
	var vargs = make([]interface{}, narg)

	for i, _ := range vargs {
		vargs[i] = new(int)
	}

	if n, err := fmt.Scanln(vargs...); n != narg-1 && err != nil {
		panic(err)
	}

	var s = make([]int, 0, len(vargs))

	for _, n := range vargs {
		s = append(s, *(n.(*int)))
	}
	sort.Ints(s)
	fmt.Println(LonelyInt(s))

}

func LonelyInt(S []int) int {
	unique := true
	if len(S) == 1 {
		return S[0]
	}
	for n := 0; n < len(S); n++ {
		i, j := S[n], S[n+1]
		if i == j {
			unique = false
			continue
		} else {
			if !unique {
				unique = true
				if j == S[len(S)-1] {
					return j
				}
				continue
			}
			return i
		}

	}
	return -1
}
