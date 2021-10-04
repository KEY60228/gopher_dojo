/*
【TRY】組み込み型（真偽値）
次のプログラムでtrueとなるケースを考えてください
trueと表示される場合のa,b,cの値にはどのようなパターンがあるか
真理値表を作成し埋めて下さい

package main
func main() {
	var a, b, c bool
	if a && b || !c {
		println("true")
	} else {
		println("false")
	}
}
*/
package main

func main() {
	var a, b, c bool
	if a && b || !c {
		println("true")
	} else {
		println("false")
	}
}
