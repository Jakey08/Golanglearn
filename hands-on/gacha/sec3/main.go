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

//added tickets
type player struct {
	tickets int
}





func main() {
	rand.Seed(time.Now().Unix())

	p := player{tickets:10}

	n := inputN(&p)
	results, summary := drawN(&p, n) //substitute multiple results

	fmt.Println(results)
	fmt.Println(summary)
}

//I created inputN_func which is to return n
func inputN(p *player) int {
	var n int
	for {
		fmt.Print("ガチャを引く<回数>")
		fmt.Scanln(&n)
		if n > 0 && n <= p.tickets {
			break
		}
		fmt.Printf("1以上%d以下の数を入力してください\n", p.tickets)
	}
	return n
}

//func f_name(arg_name type, arg_name2 type) (return_value1, return_value2){}
//ex) func swap(x, y int)(int, int){}
//mapping key and value
func drawN(p *player, n int) ([]card, map[rarity]int) {
	p.tickets-=n
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
