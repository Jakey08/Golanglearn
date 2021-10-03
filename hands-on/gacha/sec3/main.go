package main

import (
	"fmt"
	"math/rand"
	"time"
)

type rarity string

const (
	rarityN  rarity = "ノーマル"
	rarityR  rarity = "R"
	raritySR rarity = "SR"
	rarityXR rarity = "XR"
)

type card struct {
	rarity rarity
	name   string
}

func main() {
	rand.Seed(time.Now().Unix())

	n := inputN()
	results, summary := drawN(n) //substitute multiple results

	fmt.Println(results)
	fmt.Println(summary)
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

//func f_name(arg_name type, arg_name2 type) type {}
func drawN(n int) ([]card, map[rarity]int) {
	results := make([]card, n)
	summary := make(map[rarity]int)
	//return summary as each result of rarities

	for i := 0; i < n; i++ {

		results[i] = draw() //not draw(n)
		summary[results[i].rarity]++
	}
	return results, summary
}

//func func_name() type(*this is return value) {}
//ex) func hello() string{}
func draw() card {
	num := rand.Intn(100)

	switch {
	case num < 80:
		return card{rarity: rarityN, name: "スライム"}
	case num < 95:
		return card{rarity: rarityR, name: "オーク"}
	case num < 99:
		return card{rarity: raritySR, name: "ドラゴン"}
	default:
		return card{rarity: rarityXR, name: "イフリート"}
	}
}
