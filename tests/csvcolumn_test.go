package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// "Kliendi konto";"Reatüüp";"Kuupäev";"Saaja/Maksja";"Selgitus";"Summa";"Valuuta";"Deebet/Kreedit";"Arhiveerimistunnus";"Tehingu tüüp";"Viitenumber";"Dokumendi number";

func TestFormat(t *testing.T) {

	// replace the value of this variable to test your csv
	path_to_your_csv_statement := "C:/Users/rober/Downloads/statement.csv"

	// expected format of the csv
	format := [13]string{"Kliendi konto", "Reatüüp", "Kuupäev", "Saaja/Maksja", "Selgitus", "Summa", "Valuuta", "Deebet/Kreedit", "Arhiveerimistunnus", "Tehingu tüüp", "Viitenumber", "Dokumendi number"}

	file, err := os.Open(path_to_your_csv_statement)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	firstRecord, err := reader.Read()
	if err != nil {
		panic(err)
	}

	if assert.Equal(t, format[:], firstRecord) {
		fmt.Println("\033[32mHeaders OK\033[0m")
	}
	fmt.Println("\033[32mTEST PASSED\033[0m")
}
