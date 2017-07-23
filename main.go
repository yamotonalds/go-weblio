package main

import (
	"fmt"
	"io"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	inputWords := readWords()
	translatedWords := translateWords(inputWords)

	for translated := range translatedWords {
		fmt.Println(translated)
	}
}

func readWords() chan string {
	inputWords := make(chan string, 10)

	go func() {
		for {
			var word string
			_, err := fmt.Scanln(&word)
			if err != nil {
				if err != io.EOF {
					fmt.Fprintf(os.Stderr, "単語が読み込めませんでした: %v\n", err)
				}
				break
			}

			inputWords <- word
		}
		close(inputWords)
	}()

	return inputWords
}

func translateWords(inputWords chan string) chan string {
	translatedWords := make(chan string, 10)

	go func() {
		for word := range inputWords {
			translated, err := translate(word)
			if err != nil {
				fmt.Fprintf(os.Stderr, "翻訳エラー: %v\n", err)
				continue
			}
			translatedWords <- translated
		}
		close(translatedWords)
	}()

	return translatedWords
}

func translate(word string) (string, error) {
	doc, err := goquery.NewDocument("http://ejje.weblio.jp/content/" + word)
	if err != nil {
		return "", fmt.Errorf("%s: 翻訳ページが読み込めませんでした(%v)", word, err)
	}

	contentExplanation := doc.Find(".content-explanation").First()
	if contentExplanation == nil {
		return "", fmt.Errorf("%s: 翻訳が見つかりませんでした", word)
	}

	return contentExplanation.Text(), nil
}
