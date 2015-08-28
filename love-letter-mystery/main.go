package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var s []byte
		fmt.Scan(&s)
		var cost int
		for i, _ := range s[:(len(s)/2)+1] {
			lchar, rchar := s[i], s[(len(s)-1)-i]
			var lcost, rcost int
			if lchar < rchar {
				rcost += int(rchar - lchar)
				s[(len(s)-1)-i] -= byte(rcost)
			} else if lchar > rchar {
				lcost += int(lchar - rchar)
				s[i] -= byte(lcost)
			} else {
				continue
			}
			cost += rcost + lcost
		}
		fmt.Println(cost)

	}

}
