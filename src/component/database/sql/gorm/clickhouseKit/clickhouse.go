package postgresqlKit

import (
	"github.com/richelieu042/chimera/v3/src/component/database/sql/gorm/gormKit"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

// NewGormDB
/*
@param dsn e.g."clickhouse://gorm:gorm@localhost:9942/gorm?dial_timeout=10s&read_timeout=20s"
*/
func NewGormDB(dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	dialector := clickhouse.Open(dsn)
	return gormKit.NewDB(dialector, opts...)
}
