package database

import "gorm.io/gorm"

// Fetch all logs matching address and topic0 from timestamp range [from, to), order by timestamp
func FetchLogsByAddressAndTopic0(db *gorm.DB, address string, topic0 string,
	from int64, to int64) ([]Log, error) {
	var logs []Log
	err := db.Where("address = ? AND topic0 = ? AND timestamp >= ? AND timestamp < ?",
		address, topic0, from, to).Order("timestamp").Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}
