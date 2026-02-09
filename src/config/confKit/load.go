package confKit

//import (
//	"github.com/richelieu042/chimera/v3/src/file/fileKit"
//	"github.com/richelieu042/chimera/v3/src/log/console"
//	"github.com/zeromicro/go-zero/core/conf"
//)
//
//// MustLoad 读取配置文件并反序列化为 指定结构体指针ptr.
///*
//https://github.com/zeromicro/go-zero/tree/master/core/conf
//go-zero/core/conf 的缺陷: 不能自动将string类型转换为int类型，会直接返回error.
//
//PS:
//(0) 不支持 .properties，字段验证相关还是用 validateKit 吧
//(1) default对应的是 该行被注释掉（而非值为空） 的情况;
//(2) default和 '|' 一起使用;
//(3) range和 ':' 一起使用.
//
//!!!:
//(0) 所有字段首字母大写 && 有json tag;
//(1) 结构体可以参考 go-zero的 rest.RestConf，可以通过 tag 控制配置的值（默认值default、范围range、可选optional...；可以组合）
//(2) 值为数组时要注意（要么有值(size > 0)，要么全注释掉），以免返回error.
//(3) tag range(定义当前字段数值范围)支持:
//	[:5] (:5] [:5) (:5)
//	[1:] [1:) (1:] (1:)
//	[1:5] [1:5) (1:5] (1:5)
//
//	e.g.
//	type config struct {
//		Number1 int           `json:"number1,range=[1:]"`
//	}
//
//	可能返回的error:
//	(a) wrong number range setting: 不满足tag range
//
//@param path	配置文件的路径（推荐使用.yaml）
//@param ptr	[不能为nil] 结构体实例的指针
//@param opts e.g. conf.UseEnv()
//
//e.g.	结构体属性的tag（https://www.w3cschool.cn/gozero/gozero-eo623nm5.html）
//	json:"name,optional"
//	json:"gender,options=male|female"
//	json:"gender,default=male"
//	json:"age,range=[0:120]"
//
//e.g.1	组合多个tag
//	Port     int    `json:"port,optional,range=[-1:65535]"`
//	Port     int    `json:"port,default=-1,range=[-1:65535]"`
//*/
//func MustLoad(path string, ptr any, options ...conf.Option) {
//	if err := Load(path, ptr, options...); err != nil {
//		console.Fatalf("Fail to load, error: %s", err.Error())
//	}
//}
//
//// Load
///*
//PS: 不支持 .properties，字段验证相关还是用 validateKit 吧.
//*/
//func Load(path string, ptr any, options ...conf.Option) error {
//	if err := fileKit.AssertExistAndIsFile(path); err != nil {
//		return err
//	}
//
//	return conf.Load(path, ptr, options...)
//}
