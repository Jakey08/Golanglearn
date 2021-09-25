package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix()) //rand.Seed is not included in 'for loop'.

	for i := 1; i < 12; i++ {

		num := rand.Intn(100)
		fmt.Printf("%d回目", i)

		switch {
		case num < 80:
			fmt.Println("ノーマル")
		case num < 94:
			fmt.Println("R")
		case num < 99:
			fmt.Println("SR")
		default:
			fmt.Println("XR")
		}
	}

}
