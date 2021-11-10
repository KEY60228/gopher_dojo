/*
【TRY】単体テストを行おう
偶数判定関数のテスト
	偶数が与えられた場合にtrueを返す関数を考える
	テーブル駆動テストを行う
	テストケースとしてふさわしいものを考えテストを書こう
	各テストケースをサブテストとして実行しよう
	coverprofileをみてみよう
*/
package mypkg

import "fmt"

func main() {
	fmt.Println(isEven(2))
}

func IsEven(n int) bool {
	return n%2 == 0
}
