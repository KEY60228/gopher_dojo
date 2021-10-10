/*
【TRY】レシーバに変更を与える
次のプログラムを正しく動作するようにしてください
Incメソッドは自身を1ずつ加算する
今の実装だと正しく動かない
動かない理由を考え、意図通り動くように修正してください

package main
type MyInt int
func (n MyInt) Inc() { n++ }
func main() {
	var n MyInt
	println(n)
	n.Inc()
	println(n)
}
*/
package main

import "fmt"

type MyInt int

func (n *MyInt) Inc() {
	*n++
}

func main() {
	var n MyInt
	fmt.Println(n)
	(&n).Inc()
	fmt.Println(n)
}
