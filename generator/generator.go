package generator

import (
	"math/rand"
	"time"
)

type Generator interface {
	GenerateRandomNumbers() []int
}

type RandomNumbersGenerator struct {
	min, max, len int
}

func NewRandomNumbersGenerator(min, max, len int) Generator {
	rand.Seed(time.Now().UnixNano())

	return &RandomNumbersGenerator{
		min: min,
		max: max,
		len: len,
	}
}

func (r *RandomNumbersGenerator) GenerateRandomNumbers() []int {
	randNums := make([]int, r.len)

	for i := range randNums {
		randNums[i] = rand.Intn(r.max-r.min+1) + r.min
	}

	return randNums
}
