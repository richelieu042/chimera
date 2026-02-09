package mysqlKit

import (
	"github.com/richelieu042/chimera/v3/src/component/database/sql/gorm/gormKit"
	"github.com/richelieu042/chimera/v3/src/core/interfaceKit"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewGormDB
/*
PS:
(1) 参考: https://gorm.io/zh_CN/docs/connecting_to_the_database.html#MySQL
(2) 个人实测，如果密码包含特殊字符（e.g.@或/）或者中文，并不需要进行编码（e.g.Golang中使用 url.QueryEscape()）.

@param dsn	"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
			查询参数:
			(1) charset=utf8mb4:	指定字符集为 utf8mb4，适用于支持表情符号等扩展字符集的场景
			(2) parseTime=True: 	表示将数据库中的时间值解析为 Go 的 time.Time 类型
			(3) loc=Local:	 		指定时区为本地时区，适用于数据库时间和本地时间保持一致的情况
@param opts	可以不传
*/
func NewGormDB(dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	dialector := mysql.Open(dsn)
	return gormKit.NewDB(dialector, opts...)
}

// NewGormDBOptionally 自定义驱动.
/*
参考: https://gorm.io/zh_CN/docs/connecting_to_the_database.html#MySQL
*/
func NewGormDBOptionally(mysqlConfig *mysql.Config, opts ...gorm.Option) (*gorm.DB, error) {
	if err := interfaceKit.AssertNotNil(mysqlConfig, "mysqlConfig"); err != nil {
		return nil, err
	}

	dialector := mysql.New(*mysqlConfig)
	return gormKit.NewDB(dialector, opts...)
}
