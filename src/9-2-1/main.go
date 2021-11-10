/*
【TRY】チャンネルを使ってみよう
次のコードのTODOを実装してみよう
	標準入力から受け取った文字列を出力するコードになります
func input(r io.Reader) <-chan string {
	// TODO: チャネルを作る
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			// TODO: チャネルに読み込んだ文字列を送る
		}
		// TODO: チャネルを閉じる
	}()
	// TODO: チャネルを返す
}
func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print(">")
		fmt.Println(<-ch)
	}
}
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print(">")
		fmt.Println(<-ch)
	}
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}
