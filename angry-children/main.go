package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	var nset = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nset[i])
	}
	if !sort.IntsAreSorted(nset) {
		sort.Ints(nset)
	}
	var fairness = math.Inf(0)
	for i := 0; i+k <= n; i++ {
		kset := nset[i : i+k]
		f := float64(max(kset) - min(kset))
		if f < fairness {
			fairness = f
		}
	}
	fmt.Println(int(fairness))
}

func min(kset []int) int {
	//fmt.Printf("min(%v)=>%v\n", kset, kset[0])
	return kset[0]

}
func max(kset []int) int {
	//fmt.Printf("max(%v)=>%v\n", kset, kset[len(kset)-1])
	return kset[len(kset)-1]
}
