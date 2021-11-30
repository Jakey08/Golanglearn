package gachapack

type Rarity string //パッケージ外へのエクスポートは先頭が大文字になる
		   //他のパッケージから利用できるようになる

const (
rarityN  rarity = "N"
rarityR  rarity = "R"
raritySR rarity = "SR"
rarityXR rarity = "XR"
)

func (r rarity) String() string {
	return string(r)
}

//TODO フィールドのエクスポートをする

type Card struct {
	Rarity Rarity // レア度
	Name   string // 名前
}

func (c *card) String() string {

	return c.Rarity.String() + ":" + c.name
}
