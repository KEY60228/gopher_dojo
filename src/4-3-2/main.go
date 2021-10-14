/*
【TRY】modulesを使ってみよう
modules用のプロジェクトを作成する

github.com/tenntenn/greetingのv1.0.0を使ってみる
fmt.Println(greeting.Do())をやってみる
-> 4-3-1

github.com/tenntenn/greetingのv2.0.0を使ってみる
fmt.Println(greeting.Do(time.Now()))をやってみる
*/
package main

import (
	"fmt"
	"time"

	greeting "github.com/tenntenn/greeting/v2"
)

func main() {
	fmt.Println(greeting.Do(time.Now()))
}
