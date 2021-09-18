package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	for i := 1; i < 4; i++ {
		fmt.Printf("%d回目おみくじ結果発表 : ", i)

		number := rand.Intn(8)

		switch number {
		case 0:
			fmt.Println("大凶です 明日は仕事休んだほうがいいでしょう\n")
		case 1:
			fmt.Println("ちょっと凶です カフェにでも行きましょう\n")
		case 2:
			fmt.Println("吉です 特に言うことはないです\n")
		case 3:
			fmt.Println("まあまあ吉です たまには朝ごはん作ってみては？\n")
		case 4:
			fmt.Println("小吉です 明日は仕事休んだほうがいいでしょう\n")
		case 5:
			fmt.Println("中吉です 特に言うことはないです\n")
		case 6:
			fmt.Println("大吉です 明日は仕事休んだほうがいいでしょう\n")
		case 7:
			fmt.Println("超吉です なんかいいこと起こるかも？\n")
		}

	}

}
