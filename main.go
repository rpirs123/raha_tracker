package main

import (
	"fmt"
	"log"
	"os"
	util "raha_tracker/src"
	"raha_tracker/src/cli"
	"raha_tracker/src/iniutil"
	"raha_tracker/src/output"
	csvparser "raha_tracker/src/parser"
	"strconv"
	"strings"
)

var (
	kulud map[string]float64
	tulud map[string]float64
)

func main() {
	csvFile := os.Args[1]
	kulud = make(map[string]float64)
	tulud = make(map[string]float64)

	if _, err := os.Stat("config"); os.IsNotExist(err) {
		fmt.Println("Config file inside root not found...Creating...")
		iniutil.CreateConfig()
		fmt.Println("Config file created successfully")
	}

	iniutil.LoadConfig()

	records, err := csvparser.ParseCsv(csvFile)
	if err != nil {
		log.Fatal(err)
	}

	recLength := len(records)
	for i := range records {
		// skip first two records
		if i == 0 || i == 1 {
			continue
		}

		if i >= recLength-3 {
			continue
		}
		rec := records[i]

		rcrd := util.ParseRecord(rec)

		if rcrd.Supplier != "" {
			cli.RunPrompts(&rcrd)
		}

		addRecord(&rcrd)
	}
	output.ShowOutput(kulud, tulud)
}

func addRecord(rec *util.Record) {

	clean := strings.ReplaceAll(rec.Amount, ",", ".")
	amount, err := strconv.ParseFloat(clean, 64)
	if err != nil {
		// handle error
		fmt.Println("Error during parsing string to float:", err)
	}

	switch rec.TransactionType {
	case "D":
		kulud[rec.Category] += amount
	case "K":
		tulud[rec.Category] += amount
	}

}
