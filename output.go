package main

import (
	"fmt"
	"os"
	"strings"
)

func SaveToFile(filename string, nums []int) error {
	data := strings.Trim(strings.Replace(fmt.Sprint(nums), " ", ", ", -1), "[]")
	return os.WriteFile(filename, []byte(data), 0644)
}
