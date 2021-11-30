package gacha

import "fmt"

type player struct {
	tickets int // ガチャ券の枚数
	coin    int // コイン
}

// TODO: 引数にガチャ券とコインの枚数をもらい、
// それぞれをフィールドに設定したPlayer型の値を生成し、
// そのポインタを返すNewPlayer関数を作る

func NewPlayer(int *player.tickets, int *player.coin) {
	return 
}

// プレイヤーが行えるガチャの回数
func (p *player) drawableNum() int {

// TODO: ガチャが行える回数を返す
// ガチャ券は1枚で1回、コインは10枚で1回ガチャが行える
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
	p.coin -= n * 10 // 1回あたり10枚消費する
}