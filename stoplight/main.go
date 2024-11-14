package main

import (
	"fmt"
	"time"
)

func main() {
	stage := 0
	
	for {
		switch stage {
        case 0:
            fmt.Println("Red")
			time.Sleep(60 * time.Second)
        case 1:
            fmt.Println("Yellow")
			time.Sleep(15 * time.Second)
        case 2:
            fmt.Println("Green")
			time.Sleep(120 * time.Second)
        }
		
		if stage == 2 {
			stage = 0
		} else {
			stage++
		}
	}
}