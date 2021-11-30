package gachapack

import 	(
	"math/rand"
	"time"

)
	

func init(){
	//乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())
}

func DrawN(p *player, n int) ([]*card, map[Rarity]int) {
	p.draw(n)

	results := make([]*card, n)
	summary := make(map[Rarity]int)
	for i := 0; i < n; i++ {
		results[i] = draw()
		summary[results[i].rarity]++
	}

	// 変数resultsとsummaryの値を戻り値として返す
	return results, summary
}

func draw() *Card {
	num := rand.Intn(100)

	switch {
	case num < 80:
		return &card{Rarity: RarityN, Name: "スライム"}
	case num < 95:
		return &card{Rarity: RarityR, Name: "オーク"}
	case num < 99:
		return &card{Rarity: RaritySR, Name: "ドラゴン"}
	default:
		return &card{Rarity: RarityXR, Name: "イフリート"}
	}
}
