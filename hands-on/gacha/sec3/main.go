package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	n := inputN()

	for i := 1; i <= n; i++ {
		draw()
	}
}

//I created inputN_func which is to return n
func inputN() int {
	var n int
	for {
		fmt.Print("ガチャを引く<回数>")
		fmt.Scanln(&n)
		if n > 0 {
			break
		}
		fmt.Println("もう一度入力してくだだい")
	}
	return n
}

func draw() {
	num := rand.Intn(100)

	switch {
	case num < 80:
		fmt.Println("ノーマル")
	case num < 95:
		fmt.Println("R")
	case num < 99:
		fmt.Println("SR")
	default:
		fmt.Println("XR")
	}
}
