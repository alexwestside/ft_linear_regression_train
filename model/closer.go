package model

import (
	"io"
	"log"
)

func (l *Lnreg) Closer(w io.Closer) {
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}
}
