/*
【TRY】インタフェースを作ろう
Stringerインタフェース
	String() string メソッドを持つインタフェースを作る
	そして3つ以上Stringerインタフェースを実装する型を作る
インタフェースを受け取る関数
	Stringerインタフェースを引数で受け取る関数を作る
	受け取った値を上記の3つの具象型によって分岐し、具象型の型名と値を表示する
*/
package main

import "fmt"

type Stringer interface {
	String() string
}

type S1 int
type S2 bool
type S3 float64

func (s S1) String() string {
	return fmt.Sprintf("%v", int(s))
}

func (s S2) String() string {
	return fmt.Sprintf("%v", bool(s))
}

func (s S3) String() string {
	return fmt.Sprintf("%v", float64(s))
}

func main() {
	var s1 Stringer = S1(100)
	var s2 Stringer = S2(true)
	var s3 Stringer = S3(24.12)

	fmt.Println(dp(s1))
	fmt.Println(dp(s2))
	fmt.Println(dp(s3))
}

func dp(s Stringer) string {
	return fmt.Sprintf("%T %v", s, s)
}
