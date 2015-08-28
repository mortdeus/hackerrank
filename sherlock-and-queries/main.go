package main

import (
	"fmt"
)

/*
for i = 1 to M do
    for j = 1 to N do
        if j % B[i] == 0 then
            A[j] = A[j] * C[i]
        endif
    end do
end do
*/

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var a, b, c []int

	for i := 1; i <= n; i++ {
		var aval int
		fmt.Scan(&aval)
		a = append(a, aval)
	}
	fmt.Println(a, n)

	for i := 1; i <= m; i++ {
		var bval int
		fmt.Scan(&bval)
		b = append(b, bval)
	}
	fmt.Println(b, m)
	for i := 1; i <= m; i++ {
		var cval int
		fmt.Scan(&cval)
		c = append(c, cval)
	}
	fmt.Println(c, m)

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Println(j % b[i])
			if j%b[i-1] == 0 {
				fmt.Printf("a[%v],c[%v] = %v,%v\na[j]*c[i] = %d\n\n", j, i, a[j], c[i], a[j]*c[i])
				a[j] = a[j] * c[i]
			}
		}
	}
	fmt.Println(a)
	return
}
