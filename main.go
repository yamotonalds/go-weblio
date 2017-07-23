package main

import (
	"fmt"
	"io"
	"os"
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

		fmt.Println(word)
	}
}
