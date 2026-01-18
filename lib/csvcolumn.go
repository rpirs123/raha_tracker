package util

type csvColumn struct {
	Record             []string
	Konto              int
	Reatüüp            int
	Kuupäev            int
	Maksja             int
	Selgitus           int
	Summa              int
	Valuuta            int
	Krediit            int
	Arhiveerimistunnus int
	Tehningutüüp       int
	Viitenumber        int
	Dokumendinr        int
}

var csvFormat = csvColumn{
	Konto:              0,
	Reatüüp:            1,
	Kuupäev:            2,
	Maksja:             3,
	Selgitus:           4,
	Summa:              5,
	Valuuta:            6,
	Krediit:            7,
	Arhiveerimistunnus: 8,
	Tehningutüüp:       9,
	Viitenumber:        10,
	Dokumendinr:        11,
}

func getTransactionType(record []string) string {

	switch record[csvFormat.Krediit] {
	case "K":
		return "K"
	case "D":
		return "D"
	}
	return ""
}

func getSupplier(record []string) string {
	return record[csvFormat.Maksja]
}

func getTimestamp(record []string) string {
	return record[csvFormat.Kuupäev]
}

func getCurrency(record []string) string {
	return record[csvFormat.Valuuta]
}

func getAmount(record []string) string {
	return record[csvFormat.Summa]
}
