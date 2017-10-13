package warehouse

import "time"

type Warehouse interface {
	ExportTableSchema() Schema
	LastSyncPoint() (time.Time, error)
	SaveSyncPoint(id int, stop time.Time) error
	LoadToWarehouse(filename string) error
	ValueToString(val interface{}, f Field) string
}