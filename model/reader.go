package model

import (
	"os"
	"encoding/csv"
)

func (l *Lnreg) Reader() ([][]string, error) {

	file, err := os.Open(l.File)
	if err != nil {
		return nil, err
	}
	defer l.Closer(file)

	data := csv.NewReader(file)

	records, err := data.ReadAll()
	if err != nil {
		return nil, err
	}

	l.Headers = records[0]

	return records[1:], nil
}
