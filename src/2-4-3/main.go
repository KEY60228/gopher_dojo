/*
【TRY】おみくじプログラムを作ろう
おみくじプログラム
サイコロを転がし、出た目によって運勢を占うプログラム
6：大吉
5,4：中吉
3,2：吉
1：凶
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)

	switch i := rand.Intn(5) + 1; i {
	case 1:
		fmt.Println("凶")
	case 2, 3:
		fmt.Println("吉")
	case 4, 5:
		fmt.Println("中吉")
	case 6:
		fmt.Println("大吉")
	}
}
