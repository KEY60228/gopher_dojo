/*
【TRY】画像変換コマンドを作ろう

次の仕様を満たすコマンドを作って下さい
ディレクトリを指定する
指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
ディレクトリ以下は再帰的に処理する
変換前と変換後の画像形式を指定できる（オプション）

以下を満たすように開発してください
mainパッケージと分離する
自作パッケージと標準パッケージと準標準パッケージのみ使う
準標準パッケージ：golang.org/x以下のパッケージ
ユーザ定義型を作ってみる
GoDocを生成してみる
Go Modulesを使ってみる
*/
package main

import (
	"flag"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/fs"
	"log"
	"path/filepath"

	"dojo/5-3-2/converter"
)

var bt = flag.String("bt", "jpeg", "filetype before conversion")
var at = flag.String("at", "png", "filetype after conversion")

func main() {
	flag.Parse()

	result, err := converter.IsValidExt(*bt)
	if err != nil {
		log.Fatal(err)
	}
	if !result {
		log.Fatal("対応する拡張子はjpg, png, gifのみです")
	}

	result, err = converter.IsValidExt(*at)
	if err != nil {
		log.Fatal(err)
	}
	if !result {
		log.Fatal("対応する拡張子はjpg, png, gifのみです")
	}

	err = filepath.Walk(flag.Arg(0), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		var ext string
		if filepath.Ext(info.Name()) == ".jpg" {
			ext = "jpeg"
		} else {
			ext = filepath.Ext(info.Name())[1:]
		}

		result, err := converter.IsValidExt(ext)
		if err != nil {
			return err
		}
		if !result || ext != *bt {
			return nil
		}

		err = converter.Do(path, *at)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
