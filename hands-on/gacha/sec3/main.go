package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	draw()
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
