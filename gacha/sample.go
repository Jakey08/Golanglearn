//using slice p33

package main

//import "fmt"

func main() {

	ns := []int{19, 86, 1, 12}
	sum := 0

	for i := 0; i < len(ns); i++ {
		sum += ns[i]

	}

	println(sum)
}


