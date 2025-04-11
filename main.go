package main

import (
	"fmt"
)

func main() {
	input := NewCLIInputReceiver()

	min, max, len, err := input.Receive()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	gen := NewRandomNumbersGenerator(min, max, len)
	nums := gen.GenerateRandomNumbers()

	err = SaveToFile("output.txt", nums)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Generated numbers:", nums)
	fmt.Println("Numbers saved to output.txt")
}
