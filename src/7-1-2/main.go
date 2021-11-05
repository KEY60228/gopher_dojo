/*
【TRY】エラー処理をまとめる
1コードポイント(rune)ずつ読み込むScannerを作れ
	初期化時にio.Readerを渡す
	bufio.Scannerと似た感じに使えること
	エラー処理を纏める
*/
package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"unicode/utf8"
)

type RuneScanner struct {
	r   io.Reader
	buf [16]byte
}

func (s *RuneScanner) Scan() (rune, error) {
	n, err := s.r.Read(s.buf[:])
	if err != nil {
		return 0, err
	}

	r, size := utf8.DecodeRune(s.buf[:n])
	if r == utf8.RuneError {
		return 0, errors.New("RuneError")
	}

	s.r = io.MultiReader(bytes.NewReader(s.buf[size:n]), s.r)
	return r, nil
}

func main() {
	s := NewRuneScanner(strings.NewReader("Hello, World!"))

	for {
		r, err := s.Scan()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%c\n", r)
	}
}

func NewRuneScanner(r io.Reader) *RuneScanner {
	return &RuneScanner{r: r}
}
