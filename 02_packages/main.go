package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/valeriatisch/golang-workshop/02_packages/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello"))
	fmt.Println(cmp.Diff("Hello", "Hello"))
	//fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	
}
