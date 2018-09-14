package model

import (
	"os"
	"encoding/csv"
)

type Reader interface {
	Read() ([][]string, error)
}

func (l *Model) Read(path string) ([][]string, error) {

	file, err := os.Open(path)
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
