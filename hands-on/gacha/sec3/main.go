package main

import (
	"fmt"
	"math/rand"
	"time"
)

//define the type of 'rarity' which is new type of string.
type rarity string

const (
	rarityN  rarity = "ノーマル"
	rarityR  rarity = "R"
	raritySR rarity = "SR"
	rarityXR rarity = "XR"
)

//'rarity' is the type of string which is defined eariler

func (r rarity) String() string {
	return string(r)
}

type card struct {
	rarity rarity
	name   string
}

func (c *card) String() string {
	return c.rarity.String() + ":" + c.name
}

//added tickets
type player struct {
	tickets int
	coin    int
}

func (p *player) drawableNum() int {
	return p.tickets + p.coin/10

}

func (p *player) draw(n int) {
	if p.drawableNum() < n {
		fmt.Println("ガチャ券またはコインが不足しています")
		return
	}
	if p.tickets > n {
		p.tickets -= n
		return
	}
	p.tickets = 0
	p.coin -= n * 10
}

func main() {
	rand.Seed(time.Now().Unix())

	p := player{tickets: 10, coin: 100} //initialize the tickets as 10

	n := inputN(&p)
	results, summary := drawN(&p, n) //substitute multiple results

	fmt.Println(results)
	fmt.Println(summary)
}

//I created inputN_func which is to return n
//this 'int' is the type of return value.
func inputN(p *player) int {

	max := p.drawableNum()
	fmt.Printf("ガチャを引く回数を入力してください(最大:%d回)\n", max)

	var n int
	for {
		fmt.Print("ガチャを引く<回数>")
		fmt.Scanln(&n)
		if n > 0 && n <= max {
			break
		}
		fmt.Printf("1以上%d以下の数を入力してください\n", p.tickets)
	}
	return n
}

//func f_name(arg_name type, arg_name2 type) (return_value1, return_value2){}
//ex) func swap(x, y int)(int, int){}
//mapping key and value
func drawN(p *player, n int) ([]*card, map[rarity]int) {
	p.draw(n)

	results := make([]*card, n)
	summary := make(map[rarity]int)
	//return summary as each result of rarities
	for i := 0; i < n; i++ {
		results[i] = draw() //not draw(n)
		summary[results[i].rarity]++
	}
	return results, summary
}

//func func_name() type(return value) {}
//ex) func hello() string{}
func draw() *card {
	num := rand.Intn(100)

	switch {
	case num < 80:
		return &card{rarity: rarityN, name: "スライム"}
	case num < 95:
		return &card{rarity: rarityR, name: "オーク"}
	case num < 99:
		return &card{rarity: raritySR, name: "ドラゴン"}
	default:
		return &card{rarity: rarityXR, name: "イフリート"}
	}
}
