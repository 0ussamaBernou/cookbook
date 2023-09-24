package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

//	func rot13(b byte) byte {
//		switch {
//		case b <= 'Z' && b >= 'A':
//			b = 'A' + (b+13)%26
//		case b <= 'z' && b >= 'a':
//			b = 'a' + (b+13)%26
//		default:
//
//		}
//		return b
//	}

func rot13(x byte) byte {
	capital := x >= 'A' && x <= 'Z'
	if !capital && (x < 'a' || x > 'z') {
		return x // Not a letter
	}

	x += 13
	if capital && x > 'Z' || !capital && x > 'z' {
		x -= 26
	}
	return x
}

func (rot *rot13Reader) Read(stream []byte) (int, error) {
	n, err := rot.r.Read(stream)
	if err != nil {
		return 0, nil
	}
	for i, value := range stream {
		stream[i] = rot13(value)
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

}
