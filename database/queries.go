package database

import (
	"strings"

	"gorm.io/gorm"
)

// Fetch all logs matching address and topic0 from timestamp range (from, to], order by timestamp
func FetchLogsByAddressAndTopic0(db *gorm.DB, address string, topic0 string,
	from int64, to int64) ([]Log, error) {
	var logs []Log
	err := db.Where(
		"address = ? AND topic0 = ? AND timestamp > ? AND timestamp <= ?",
		strings.ToLower(strings.TrimPrefix(address, "0x")),
		strings.ToLower(strings.TrimPrefix(topic0, "0x")),
		from, to,
	).Order("timestamp").Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}

// Fetch all transactions matching toAddress and functionSig from timestamp range (from, to], order by timestamp
func FetchTransactionsByAddressAndSelector(db *gorm.DB, toAddress string, functionSig string,
	from int64, to int64) ([]Transaction, error) {
	var transactions []Transaction
	err := db.Where(
		"to_address = ? AND function_sig = ? AND timestamp > ? AND timestamp <= ?",
		strings.ToLower(strings.TrimPrefix(toAddress, "0x")),
		strings.ToLower(strings.TrimPrefix(functionSig, "0x")),
		from, to,
	).Order("timestamp").Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
