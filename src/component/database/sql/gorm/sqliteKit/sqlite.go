package sqliteKit

import (
	"github.com/richelieu042/chimera/v3/src/component/database/sql/gorm/gormKit"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewGormDB
/*
@param dsn 详见: notes/Golang/database/gorm.wps
*/
func NewGormDB(dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	dialector := sqlite.Open(dsn)
	return gormKit.NewDB(dialector, opts...)
}
