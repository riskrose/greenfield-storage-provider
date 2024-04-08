package bsdb

import (
	"time"
)

// GetLatestBlockNumber get current latest block number
func (b *BsDBImpl) GetLatestBlockNumber() (uint64, error) {
	var (
		latestBlockNumber uint64
		err               error
	)
	startTime := time.Now()
	methodName := currentFunction()
	defer func() {
		if err != nil {
			MetadataDatabaseFailureMetrics(err, startTime, methodName)
		} else {
			MetadataDatabaseSuccessMetrics(startTime, methodName)
		}
	}()

	err = b.db.Table((&Epoch{}).TableName()).Select("block_height").Take(&latestBlockNumber).Error
	return latestBlockNumber, err
}
