/*
【TRY】ライブラリを取得してみよう
github.com/tenntenn/greetingをgo getする
mainパッケージから使ってみる
インポートパスを"github.com/tenntenn/greeting"にしてインポートする
greeting.Do関数を呼び出す
*/
package main

import (
	"fmt"

	"github.com/tenntenn/greeting"
)

func main() {
	fmt.Println(greeting.Do())
}
