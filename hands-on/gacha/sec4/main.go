// STEP01: ファイルを分けよう

package main

import (
	"fmt"
	"math/rand"
	"github.com/gohandson/gacha-ja/gacha"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	// TODO: ガチャ券10枚、コイン100枚を持ったプレイヤーを作る
	p := {ticket: 10, coin : 100}

	n := inputN(p)
	// TODO: gacha.DrawN関数を呼び、変数resultsとsummaryに結果を代入する
	results, summary := gacha.DrawN()
	fmt.Println(results)
	fmt.Println(summary)
}

// TODO: 引数の型をgacha.Playerのポインタにする
func inputN(p *gacha.Player) int {

	max := p.DrawableNum()
	fmt.Printf("ガチャを引く回数を入力してください（最大:%d回）\n", max)

	var n int
	for {
		fmt.Print("ガチャを引く回数>")
		fmt.Scanln(&n)
		if n > 0 && n <= max {
			break
		}
		fmt.Printf("1以上%d以下の数を入力してください\n", max)
	}

	return n
}
