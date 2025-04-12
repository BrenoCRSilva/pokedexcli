package game

import (
	"fmt"
	"time"
)

func ExploreAnimation() {
	fmt.Print("Exploring")
	for i := 0; i < 3; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println() // Move to next line when done
}

func ShakeAnimation() {
	fmt.Print("Shake")
	for i := 0; i < 3; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println() // Move to next line when done
}
