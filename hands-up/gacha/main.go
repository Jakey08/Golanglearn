package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix()) //rand.Seed is not included in 'for loop'.

	var n int // the repeating times of gacha
	fmt.Println("1: 単発ガチャ 2: 11連ガチャ")

LOOP:
	//infinite loop
	for {

		fmt.Print(">")
		var kind int
		fmt.Scanln(&kind)
		//Users can choose 1 or 2. When it's done, it'll break LOOP.
		switch kind {
		case 1: //one gacha
			n = 1
			break LOOP
		case 2: //11 times
			n = 11
			break LOOP
		default:
			fmt.Println("もう一度入力してください")
		}
	}

	for i := 1; i <= n; i++ {

		num := rand.Intn(100)
		fmt.Printf("%d回目", i)

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

}
