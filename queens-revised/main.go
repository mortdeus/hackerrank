package main

import (
	"fmt"
)

var n int = 1000

type square struct {
	x, y int

	//locality matrix
	lmat struct {
		n,1,ne,nw 
		1,s,se,sw, *square
		ne,se,e,1,
		nw,sw,1,w	
	}
	queen    *queen
	occupied bool
}

type board struct {
	grid [n][n]*square
}

type queen struct {
	sq *square
}

func main() {

}
