package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix()) //rand.Seed is not included in 'for loop'.

	var n int // the repeating times of gacha
	// 	fmt.Println("1: 単発ガチャ 2: 11連ガチャ")

	for {
		fmt.Print("何回引きますか？>")
		fmt.Scanln(&n)

		if n > 0 {
			break
		}
		fmt.Println("もう一度入力してください")
	}

	result := make(map[string]int)
	//set up the key and mapping with its value
	//setup like this 'make([]type length, capacity)'
	//Users can choose the times to do gacha.

	for i := 0; i < n; i++ {
		//the condition of for looop is i < len(result)
		num := rand.Intn(100)
		//fmt.Printf("%d回目", i)

		switch {
		case num < 80:
			result["ノーマル"]++
			//increment the value of mapping
		case num < 95:
			result["R"]++
		case num < 99:
			result["SR"]++
		default:
			result["XR"]++
		}
	}

	fmt.Println(result)

}
