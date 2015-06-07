package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = byte('A')
	}
	return len(b), nil
}

type rot13Reader struct {
	r io.Reader
}

func (t rot13Reader) Read(b []byte) (int, error) {
	n, err := t.r.Read(b)
	for i, p := range b {
		pr := rune(p)
		if unicode.IsLetter(pr) {
			a := byte('a')
			if unicode.IsUpper(pr) {
				a = byte('A')
			}
			b[i] = (p-a+13)%26 + a
		}
	}
	return n, err
}

func io_main() {
	/*r := strings.NewReader("Hello, reader !")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v err=%v b=%v\n", n, err, b)
		fmt.Printf("b[:n]=%q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
	reader.Validate(MyReader{})*/

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
