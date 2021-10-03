/*
【TRY】奇数と偶数
1-100までの数値のうち、奇数の場合は「数値-奇数」、偶数の場合は、「数値-偶数」という表示を行うプログラムを作成してください。
条件分岐にはswitchを用いてください。
*/
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch i % 2 {
		case 0:
			fmt.Printf("%d - 偶数\n", i)
		case 1:
			fmt.Printf("%d - 奇数\n", i)
		}
	}
}
