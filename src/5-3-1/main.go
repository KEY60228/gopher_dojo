/*
【TRY】catコマンドを作ろう

作成するcatコマンドの仕様
引数でファイルパスの一覧を貰い、そのファイルを与えられた順に標準出力にそのまま出力するコマンドを作ってください。
また、-nオプションを指定すると、行番号を各行につけて表示されるようにしてください。
なお、行番号はすべてのファイルで通し番号にしてください。

$ mycat -n hoge.txt fuga.txt
1: hoge
2: hoge hoge
3: fuga
4: fugafuga
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var n = flag.Bool("n", false, "行数オプション")

func main() {
	flag.Parse()
	files := flag.Args()

	count := 1
	for i := 0; i < len(files); i++ {
		err := readFile(files[i], &count)
		if err != nil {
			fmt.Fprintln(os.Stderr, "読み込みに失敗しました: ", err)
		}
	}
}

func readFile(fn string, count *int) error {
	file, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if *n {
			fmt.Print(*count, ":\t")
			*count++
		}
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
