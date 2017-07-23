package main

import (
	"fmt"
	"io"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var word string
	for {
		_, err := fmt.Scanln(&word)
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "単語が読み込めませんでした: %v\n", err)
			}
			break
		}

		doc, err := goquery.NewDocument("http://ejje.weblio.jp/content/" + word)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: 翻訳ページが読み込めませんでした(%v)\n", word, err)
		}

		contentExplanation := doc.Find(".content-explanation").First()
		if contentExplanation == nil {
			fmt.Fprintf(os.Stderr, "%s: 翻訳が見つかりませんでした", word)
		}
		translated := contentExplanation.Text()

		fmt.Println(translated)
	}
}
