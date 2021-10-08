/*
【TRY】複数戻り値の利用
値を入れ替えるswap関数を実装してください
次のコードが正しく動作するように実装してください

package main
func main() {
	n, m := swap(10, 20)
	println(n, m)
}
*/
package main

import "fmt"

func swap(n, m int) (int, int) {
	return m, n
}

func main() {
	n, m := swap(10, 20)
	fmt.Println(n, m)
}
