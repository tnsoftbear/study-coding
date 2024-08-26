package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (mr MyReader) Read(input []byte) (int, error) {
	for idx := range input {
		input[idx] = 'A'
	}
	return len(input), nil
}

func main() {
	reader.Validate(MyReader{})
}

// https://go.dev/tour/methods/22
