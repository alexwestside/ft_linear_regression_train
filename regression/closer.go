package regression

import (
	"io"
	"log"
)

func (l *Model) Closer(w io.Closer) {
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}
}
