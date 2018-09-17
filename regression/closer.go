package regression

import (
	"io"
	"log"
)

type Closer interface {
	Closer(w io.Closer)
}

func (l *Model) Closer(w io.Closer) {
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}
}
