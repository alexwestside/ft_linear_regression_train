package regression

import (
	"os"
	"encoding/csv"
	"errors"
)

func (m *Model) Read(path string) error {

	ErrorHandler(checkFile(path))

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer m.Closer(file)

	data := csv.NewReader(file)

	records, err := data.ReadAll()
	if err != nil {
		return err
	}

	m.DataFrame.Headers = records[0]
	m.DataFrame.DF = records[1:]

	return nil
}

func checkFile(path string) error {
	f, err := os.Stat(path)
	if err != nil {
		return err
	}

	if f.Size() == 0 {
		return errors.New("ERROR: file is empty")
	}

	return nil
}
