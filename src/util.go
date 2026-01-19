package util

type Record struct {
	TransactionType string
	Supplier        string
	Timestamp       string
	Amount          string
	Currency        string
	Category        string
	SectionName     string
}

func ParseRecord(record []string) Record {

	//var kulud map[string]csvColumn
	//var tulud map[string]csvColumn

	tType := getTransactionType(record)
	supplier := getSupplier(record)
	timestamp := getTimestamp(record)
	amount := getAmount(record)
	currency := getCurrency(record)

	if supplier == "" {
		return Record{}
	}

	return Record{
		TransactionType: tType,
		Supplier:        supplier,
		Timestamp:       timestamp,
		Amount:          amount,
		Currency:        currency,
	}

}
