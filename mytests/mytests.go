package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13R *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13R.r.Read(b)
	for i, v := range b[:n] {
		if v >= 'A' && v <= 'Z' {
			b[i] = (v-'A'+13)%26 + 'A'
		} else if v >= 'a' && v <= 'z' {
			b[i] = (v-'a'+13)%26 + 'a'
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{r: s}
	io.Copy(os.Stdout, &r)
}
