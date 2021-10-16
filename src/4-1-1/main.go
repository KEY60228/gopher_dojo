/*
【TRY】パッケージを分けてみよう
greetingパッケージを作ろう
	GoLandで新しいプロジェクトを作成
	greetingパッケージとし、Goファイルを1つ作る
	GoファイルにDo関数を作り”こんにちは”と返すようにする
mainパッケージから利用しよう
	mainパッケージからgreetingパッケージをインポート
	greetingパッケージDo関数を呼び出し”こんにちは”を取得
	greeting.Do()で取得できる
	取得した”こんにちは”をfmt.Println関数で出力してみる
*/
package main

import (
	"fmt"

	"dojo/4-1-1/greeting"
)

func main() {
	fmt.Println(greeting.Do())
}
