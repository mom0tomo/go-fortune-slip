package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 擬似乱数をつくる
	rand.Seed(time.Now().UnixNano())

	// fortune slipの種類
	fs := []string{"凶", "小吉", "吉", "中吉", "大吉"}

	// 乱数をn回つくる
	n := len(fs)
	i := rand.Intn(n)

	// 占いの結果
	r := fs[i]

	// 結果を出力する
	fmt.Println("今日の運勢は", r, "です！")
}
