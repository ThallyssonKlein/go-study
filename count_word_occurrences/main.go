package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main () {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n')
	words := strings.Split(input, " ")
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[strings.TrimSpace(word)] += 1
	}

	fmt.Println("Word count: ")
	for key, value := range wordCount {
		fmt.Println(key, ":", value)
	}
}