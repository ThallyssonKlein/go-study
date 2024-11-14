package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main () {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the list of numbers separated by comma: ")
	input, _ := reader.ReadString('\n')

	numbers := strings.Split(input, ",")
	floatNumbers := make([]float64, len(numbers))

	for i, number := range numbers {
        num, err := strconv.ParseFloat(strings.TrimSpace(number), 64)
        if err != nil {
            fmt.Println("Invalid input:", number)
            return
        }
        floatNumbers[i] = num
    }

	n := len(numbers)

	for i, _ := range numbers {
		swapped := false

		for j := 0; j < n -i -1; j++ {
			if floatNumbers[j] > floatNumbers[j + 1] {
				floatNumbers[j] = floatNumbers[j +1 ]
				floatNumbers[j + 1] = floatNumbers[j]
				swapped = true
			}

			if !swapped {
				break
			}
		}
	}

	fmt.Println("Sorted numbers: ", floatNumbers)
}