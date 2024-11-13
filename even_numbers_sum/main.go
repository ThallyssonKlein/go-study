package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the list of numbers separated by comma: ")
	input, _ := reader.ReadString('\n')

	numbers := strings.Split(input, ",")
	evenSum := 0

	for _, number := range numbers {
		num, err := strconv.Atoi(strings.TrimSpace(number))

		if err != nil {
			fmt.Println("Invalid input")
			os.Exit(1)
		}

		if num%2 == 0 {
			evenSum += num
		}
	}

	fmt.Println("Sum of even numbers is:", evenSum)
}