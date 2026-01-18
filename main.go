package main

import (
	"fmt"
	"log"
	"os"
	util "raha_tracker/lib"
	"raha_tracker/lib/cli"
	"raha_tracker/lib/iniutil"
	csvparser "raha_tracker/lib/parser"
	"strconv"
)

var (
	kulud map[string]float64
	tulud map[string]float64
)

func main() {
	csvFile := os.Args[1]

	iniutil.LoadConfig()

	records, err := csvparser.ParseCsv(csvFile)
	if err != nil {
		log.Fatal(err)
	}

	recLength := len(records)
	for i := range records {

		// skip first record
		if i == 0 {
			continue
		}

		if i >= recLength-3 {
			return
		}
		rec := records[i]

		rcrd := util.ParseRecord(rec)

		if rcrd.Supplier != "" {
			cli.RunPrompts(&rcrd)
		}

		addRecord(&rcrd)
	}
}

func addRecord(rec *util.Record) {
	fmt.Println("RCCC", rec.Supplier, rec.Category, rec.SectionName)

	amount, err := strconv.ParseFloat(rec.Amount, 64)
	if err != nil {
		// handle error
		fmt.Println("Error during parsing string to float:", err)
	}

	switch rec.TransactionType {
	case "D":
		kulud[rec.Category] = amount
	case "K":
		tulud[rec.Category] = amount
	}

}
