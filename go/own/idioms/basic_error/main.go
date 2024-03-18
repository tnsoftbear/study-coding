package main

import (
	"errors"
	"fmt"
)

type BasicError struct {
	Msg string
}

func (e *BasicError) Error() string {
	return e.Msg
}

func main() {
	// Создаем цепочку ошибок
	basicErr := &BasicError{Msg: "my custom error"}
	wrappedErr := fmt.Errorf("wrapped error: %w", basicErr)

	// Проверяем эквивалентность ошибок с использованием errors.Is()
	target := &BasicError{Msg: "my custom error"}
	fmt.Println("errors.Is(wrappedErr, target): ", errors.Is(wrappedErr, target)) // false
	fmt.Println("errors.Is(wrappedErr, basicErr): ", errors.Is(wrappedErr, basicErr)) // true
	fmt.Println("errors.Is(basicErr, target): ", errors.Is(basicErr, target)) // false

	// Извлекаем конкретный тип ошибки с использованием errors.As()
	var extractedErr *BasicError
	if errors.As(wrappedErr, &extractedErr) {
		fmt.Println("Extracted error message:", extractedErr.Msg) // "my custom error"
	}
}
