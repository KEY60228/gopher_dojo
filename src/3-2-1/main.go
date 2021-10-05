/*
【TRY】奇数偶数判定関数
奇数偶数判定関数を作成してください
そして、以下のプログラムの条件式の部分で使用して下さい

package main
func main() {
	for i := 1; i <= 100; i++ {
		print(i)
		if i%2 == 0 {
			println("-偶数")
		} else {
			println("-奇数")
		}
	}
}
*/
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Print(i)
		if even_odd(i) {
			fmt.Println(" - 偶数")
		} else {
			fmt.Println(" - 奇数")
		}
	}
}

func even_odd(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}
