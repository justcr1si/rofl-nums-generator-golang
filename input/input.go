package input

import (
	"errors"
	"flag"
)

type InputReceiver interface {
	Receive() (int, int, int, string, error)
}

type CLIInputReceiver struct{}

func NewCLIInputReceiver() InputReceiver {
	return &CLIInputReceiver{}
}

func (c *CLIInputReceiver) Receive() (int, int, int, string, error) {
	min := flag.Int("min", 1, "Minimum number")
	max := flag.Int("max", 10, "Maximum number")
	len := flag.Int("len", 5, "Length of the sequence")
	filename := flag.String("file", "output.txt", "Output file name")

	flag.Parse()

	if *len <= 0 {
		return 0, 0, 0, "", errors.New("length must be greater than 0")
	}

	if *max <= *min {
		return 0, 0, 0, "", errors.New("max must be greater than min")
	}

	if *max-*min+1 < *len {
		return 0, 0, 0, "", errors.New("range too small for length")
	}

	return *min, *max, *len, *filename, nil
}
