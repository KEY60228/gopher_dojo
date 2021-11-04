/*
【TRY】エラー処理をしてみよう
Stringerインタフェースに変換する関数を作ろう
	任意の値をStringer型に変換する関数
		type Stringer{String() string}
	引数にempty interfaceを取り、Stringerとエラーを返す
		func ToStringer(v interface{}) (Stringer, error)
	返すエラー型はerrorインタフェースを実装したユーザ定義型にする
	実際に呼び出してみてエラー処理をしてみよう
		エラーが発生した場合はエラーが発生した旨を表示する
*/
package main

import (
	"fmt"
	"log"
)

type Stringer interface {
	String() string
}

type InterfaceError string

func (e InterfaceError) Error() string {
	return string(e)
}

type S1 string

func (s S1) String() string {
	return string(s)
}

type S2 string

func main() {
	// v := S1("string")
	v := S2("string")
	s, err := ToStringer(v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s.String())
}

func ToStringer(v interface{}) (Stringer, error) {
	if s, ok := v.(Stringer); ok {
		return s, nil
	}
	return nil, InterfaceError("Can't cast")
}
