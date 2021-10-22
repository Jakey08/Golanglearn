//sec2 step04-5

package main

import "fmt"

type rarity string

const (
	rarityN  rarity = "ノーマル"
	rarutyR  rarity = "R"
	raritySR rarity = "SR"
	rarityXR rarity = "XR"
)

type card struct {
	rarity rarity //rarity
	name   string
}

func main() {

	slime := card{rarity: rarityN, name: "スライム"}
	fmt.Println(slime)

	// var card struct {
	// 	rarity string // レア度
	// 	name   string // 名前
	// }

	// card.rarity = "ノーマル"
	// card.name = "スライム"

	// fmt.Println(card)
	dragon := card{rarity: raritySR, name: "ドラゴン"}
	fmt.Println(dragon)
}


