// STEP01: ファイルを分けよう

package main

import (
	"fmt"
	

	"github.com/Jakey08/Golanglearn/gacha-ja/gacha"
	


	)

func main() {
	

	// TODO: ガチャ券10枚、コイン100枚を持ったプレイヤーを作る
	p := gachapack.NewPlayer(10, 100)

	n := inputN(p)
	// TODO: gacha.DrawN関数を呼び、変数resultsとsummaryに結果を代入する
	results, summary := gachapack.DrawN(p, n) //引数p, n忘れずに!
	fmt.Println(results)
	fmt.Println(summary)
}

// TODO: 引数の型をgacha.Playerのポインタにする
func inputN(p *gachapack.Player) int {

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
