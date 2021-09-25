package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	num := rand.Intn(100)

	// TODO: 変数numが0〜79のときは"ノーマル"、
	// 80〜94のときは"R"、95〜98のときは"SR"、
	// それ以外のときは"XR"と表示する

	switch {
	case num < 80:
		fmt.Println("ノーマル")
	case num < 94:
		fmt.Println("R")
	case num < 99:
		fmt.Println("SR")
	default:
		fmt.Println("XR")
	}

}

//this is commit test2021/09/25
