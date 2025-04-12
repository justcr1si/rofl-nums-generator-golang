package main

import (
	"fmt"
	"studying/generator"
	"studying/input"
	"studying/output"
)

func main() {
	input := input.NewCLIInputReceiver()

	min, max, len, filename, err := input.Receive()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	gen := generator.NewRandomNumbersGenerator(min, max, len)
	nums := gen.GenerateRandomNumbers()

	err = output.SaveToFile(filename, nums)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Generated numbers:", nums)
	fmt.Println("Numbers saved to", filename)
}
