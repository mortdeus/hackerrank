package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	runes := (strings.Split(s, ""))
	if !sort.StringsAreSorted(runes) {
		sort.Strings(runes)
	}

	lo, hi := 0, len(runes)
	mid := ((hi / 2) * (hi % 2))
	var news = make([]string, hi)

	var cur string
	for i, j := 0, 0; i+j < len(runes); {
		if cur == "\u0000" {
			break
		}
		if cur == string("") {

			cur = runes[i+j]
			i = i + j
			j = 0
		}
	jget:
		if i+j == len(runes) {
			cur = "\u0000"
			goto skip
		}
		//fmt.Println("lookup: ", cur, runes[i+j], lo, mid, hi, i, j)
		if cur == string(runes[i+j]) {
			j++
			goto jget
		} else {
			cur = string("")
		}
	skip:

		seg := runes[i : i+j]
		if len(seg)%2 == 0 {
			//fmt.Println("even: ", cur, seg, lo, mid, hi, i, j)

			hsz := (len(seg) / 2)
			nlo := (lo + hsz)
			nhi := (hi - hsz)
			//fmt.Println(news)
			for k, v := range seg[:hsz] {
				news[lo:nlo][k] = v
			}
			for l, v := range seg[hsz:] {
				news[nhi:hi][l] = v
			}
			lo = nlo
			hi = nhi

			continue
		} else if len(seg)%2 == 1 {
			//fmt.Println("odd: ", cur, seg, lo, mid, hi, i, news[mid])
			if mid != 0 {
				if news[mid] == "" {
					soff := len(seg) / 2
					slo := mid - soff
					shi := mid + soff + 1
					//fmt.Println(news)
					for m, v := range seg {
						news[slo:shi][m] = v
					}

					continue
				}
			}
			//fmt.Println(news)
			fmt.Println("NO")
			return
		}

	}
	//fmt.Println(news)
	fmt.Println("YES")
	return
}
