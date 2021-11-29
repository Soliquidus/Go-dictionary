package main

import (
	"dictionary/dictionary"
	"fmt"
	"os"
)

func main() {
	d, err := dictionary.New("./badger")
	handleErr(err)
	defer d.Close()

	err1 := d.Add("python", "An interpreted language")
	if err1 != nil {
		return
	}
	words, entries, _ := d.List()
	for _, words := range words {
		fmt.Println(entries[words])
	}

}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error: %v\n", err)
		os.Exit(1)
	}
}
