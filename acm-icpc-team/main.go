package main

import (
	"fmt"
	"index/suffixarray"
	"sort"
)

var nusers, ntopic int

func main() {

	fmt.Scanf("%v %v", &nusers, &ntopic)

	var users = make([]topic, nusers)

	for i, _ := range users {
		buf := make([]byte, ntopic)
		fmt.Scan(&buf)
		users[i] = topic(suffixarray.New(buf).Lookup([]byte("1"), -1))
		sort.Ints(users[i])
	}

	var ttab = make([]topic, 0)

	for n, usr := range users {
		for m := n + 1; m <= nusers; m++ {
			fmt.Println("\nn=", n, "m=", m)
			if m == nusers {
				if len(ttab) < len(users) {

					fmt.Println("len(ttab) == len(users-1): ", len(ttab), "==", len(users)-1)
					fmt.Printf("[pre] append(ttab, usr): %v, %v\n", ttab, usr)
					ttab = append(ttab, usr)
					fmt.Println("[post] append(ttab, usr): %v, %v\n", ttab, usr)
				}
				break
			}
			fmt.Printf("[pre] append(ttab, merge(users[n], users[m])): %v, merge(%v, %v)\n", ttab, users[n], users[m])
			ttab = append(ttab, merge(users[n], users[m]))
			fmt.Printf("[post] append(ttab, merge(users[n], users[m])): %v, merge(%v, %v)\n", ttab, users[n], users[m])
		}
	}

	fmt.Println("\n", users, "\n")
	fmt.Println("\n", ttab, "\n")

}

type topic []int

func merge(x, y topic) (z topic) {
	xsz, ysz := len(x), len(y)
	if xsz+ysz == 0 {
		return nil
	}
	if xsz == 0 && ysz > 0 {
		return y
	}
	if ysz == 0 && xsz > 0 {
		return x
	}

	var i, j = 0, 0

	z = make(topic, ntopic)
	for k, _ := range z {
		z[k] = -1

		fmt.Printf("i < len(%v) && j < len(%v): %v < %v && %v < %v\n", x, y, i, xsz, j, ysz)

		if i < len(x) && j < len(y) {
			fmt.Printf("z[%v] x[%v] y[%v]: %v %v %v\n", k, i, j, z[k], x[i], y[j])
			switch {
			case x[i] == y[j] && x[i]+y[j] > -2:
				fmt.Printf("x[i] == y[j] && x[i]+y[j] > -2: z[%v](%v) = x[%v](%v)\n", k, z[k], i, x[i])
				z[k] = x[i]
			case x[i] < y[j] && x[i] > -1:
				fmt.Printf("x[i] < y[j] && x[i] > -1: z[%v](%v) = x[%v](%v)\n", k, z[k], i, x[i])
				z[k] = x[i]
				k++
				z[k] = y[j]
			case x[i] > y[j] && y[j] > -1:
				fmt.Printf("x[i] > y[j] && y[j] > -1: z[%v](%v) = y[%v](%v)\n", k, z[k], j, y[j])
				z[k] = y[j]
				k++
				z[k] = x[i]
			}
			fmt.Printf("z[%v] x[%v] y[%v]: %v %v %v\n", k, i, j, z[k], x[i], y[j])

		} else {

			switch {
			case xsz < ysz && xsz > 0:
				fmt.Println(z[k:], x[i:])
				z[k] = x[xsz-1]

			case ysz < xsz && ysz > 0:
				fmt.Println(z[k:], y[j:])
				z[k] = y[ysz-1]

			case xsz == ysz && k < len(z):
				return z[:k+1]
				/*
					if x[xsz-1] < y[ysz-1] {
						z[k], z[k+1] = x[xsz-1], y[ysz-1]
					} else if x[xsz-1] > y[ysz-1] {
						z[k], z[k+1] = y[ysz-1], x[xsz-1]
					} else {
						z[k] = x[xsz-1]
					}
				*/
			}
		}

		switch {
		case i < xsz && j < ysz:
			i, j = i+1, j+1
		case i < xsz:
			i++
		case j < ysz:
			j++

		}

	}
	return
}
