package pgKit

import (
	"github.com/richelieu042/chimera/v3/src/component/database/sql/gorm/gormKit"
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewGormDB
/*
参考: https://gorm.io/zh_CN/docs/connecting_to_the_database.html#PostgreSQL

@param dsn e.g."host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
*/
func NewGormDB(dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	dialector := postgres.Open(dsn)
	return gormKit.NewDB(dialector, opts...)
}

// NewGormDBOptionally 自定义驱动.
/*
参考: https://gorm.io/zh_CN/docs/connecting_to_the_database.html#PostgreSQL
*/
func NewGormDBOptionally(pgConfig *postgres.Config, opts ...gorm.Option) (*gorm.DB, error) {
	if err := interfaceKit.AssertNotNil(pgConfig, "pgConfig"); err != nil {
		return nil, err
	}

	dialector := postgres.New(*pgConfig)
	return gormKit.NewDB(dialector, opts...)
}
