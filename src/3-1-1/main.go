/*
【TRY】【TRY】組み込み型（数値）
次のプログラムはコンパイルが通るでしょうか？
動作確認を行ってみてください
コンパイルが通らない場合はなぜ通らないか考え修正してください

package main
func main() {
	var sum int
	sum = 5 + 6 + 3
	avg := sum / 3
	if avg > 4.5 {
		println("good")
	}
}
*/
package main

func main() {
	var sum int
	sum = 5 + 6 + 3
	avg := float64(sum) / 3
	if avg > 4.5 {
		println("good")
	}
}
