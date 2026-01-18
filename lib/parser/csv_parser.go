package csvparser

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// currently only works for swedbank format
func ParseCsv(csvFile string) ([][]string, error) {

	fi, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(fi)
	r.Comma = ';'

	var records [][]string

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}
