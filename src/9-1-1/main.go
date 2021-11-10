/*
【TRY】ゴールーチンを使ってみよう
次のコードを実行してみよう
func main() {
	defer fmt.Println("main done")
	go func() {
		defer fmt.Println("goroutine1 done")
		time.Sleep(3 * time.Second)
	}()
	go func() {
		defer fmt.Println("goroutine2 done")
		time.Sleep(1 * time.Second)
	}()
	time.Sleep(5 * time.Second)
}
また、スリープする時間を変えてみよう
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main done")
	go func() {
		defer fmt.Println("goroutine1 done")
		time.Sleep(1 * time.Second)
	}()
	go func() {
		defer fmt.Println("goroutine2 done")
		time.Sleep(2 * time.Second)
	}()
	time.Sleep(5 * time.Second)
}
