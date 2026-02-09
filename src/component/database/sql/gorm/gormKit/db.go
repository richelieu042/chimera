package gormKit

import (
	"time"

	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"gorm.io/gorm"
)

func polyfillOpts(opts ...gorm.Option) []gorm.Option {
	if len(opts) > 0 {
		// do nothing
		return opts
	}

	opts = []gorm.Option{&gorm.Config{
		Logger: GetDefaultLogger(),
	}}
	return opts
}

// NewDB
/*
@param dialector 方言（针对不同的数据库; 不能为nil!）
*/
func NewDB(dialector gorm.Dialector, opts ...gorm.Option) (*gorm.DB, error) {
	if err := interfaceKit.AssertNotNil(dialector, "dialector"); err != nil {
		return nil, err
	}
	opts = polyfillOpts()

	db, err := gorm.Open(dialector, opts...)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	/* (1) ping */
	if err := sqlDB.Ping(); err != nil {
		return nil, errorKit.Wrapf(err, "fail to ping")
	}

	/* (2) 连接池（pool）的默认配置，后续可以按照业务需求进行更改 */
	ConfigurePoolWithSqlDB(sqlDB, 1024, 8192, time.Minute*30)

	return db, nil
}
