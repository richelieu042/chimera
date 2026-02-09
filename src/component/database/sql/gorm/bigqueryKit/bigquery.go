package postgresqlKit

import (
	"github.com/richelieu042/chimera/v3/src/component/database/sql/gorm/gormKit"
	"gorm.io/driver/bigquery"
	"gorm.io/gorm"
)

// NewGormDB
/*
@param dsn e.g."bigquery://projectid/location/dataset"
*/
func NewGormDB(dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	dialector := bigquery.Open(dsn)
	return gormKit.NewDB(dialector, opts...)
}
