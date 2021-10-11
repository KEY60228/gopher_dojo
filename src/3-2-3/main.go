/*
【TRY】ポインタ
値を入れ替えるswap2関数を実装してください
次のコードが正しく動作するように実装してください

package main
func main() {
	n, m := 10, 20
	swap2(&n, &m)
	println(n, m)
}
*/
package main

import "fmt"

func swap2(n, m *int) {
	*n, *m = *m, *n
}

func main() {
	n, m := 10, 20
	swap2(&n, &m)
	fmt.Println(n, m)
}
