package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the radius of the circle: ")
	input, _ := reader.ReadString('\n')

	radius, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Invalid input")
		os.Exit(1)
	}

	c := Circle{radius: radius}
	fmt.Println(c.Area())
}