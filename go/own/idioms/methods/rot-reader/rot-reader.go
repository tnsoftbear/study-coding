package main

import (
	"io"
	"os"
	"strings"
)

const (
	INPUT  string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	OUTPUT string = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) rot13(b byte) (byte, bool) {
	for i := range INPUT {
		if INPUT[i] == b {
			return OUTPUT[i], true
		}
	}
	return 0, false
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	var total int
	for {
		n, err := rot.r.Read(p)
		if err == io.EOF {
			return total, err // нужен возврат с ошибкой, иначе попадаем в бесконечный цикл
		}
		for i := 0; i < n; i++ {
			ch, ok := rot.rot13(p[i])
			if ok {
				p[i] = ch
			}
		}
		total += n
	}
	// return total, nil
}

// Альтернативное решение
// func (rot *rot13Reader) Read(p []byte) (int, error) {
// 	n, err := rot.r.Read(p)
// 	if err != nil {
// 		return n, err
// 	}
// 	for i := 0; i < n; i++ {
// 		ch, ok := rot.rot13(p[i])
// 		if ok {
// 			p[i] = ch
// 		}
// 	}
// 	return n, nil
// }

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
