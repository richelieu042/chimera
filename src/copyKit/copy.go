package copyKit

import (
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/gogf/gf/v2/util/gutil"
	"github.com/jinzhu/copier"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
)

// Copy 浅拷贝
/*
@param dest	(1) 如果为nil，将返回error
			(2) 必须是指针类型
@param src 	(1) 如果为nil，将返回error
			(2) 指针类型 || 结构体实例

e.g.
	b := &bean{
		Id: 666,
	}
	src := map[string]interface{}{
		"b":   false,
		"tmp": b,
	}
	dest := make(map[string]interface{})
	if err := copyKit.Copy(&dest, src); err != nil {
		panic(err)
	}

	// {"b":false,"tmp":{"Id":666}} <nil>
	fmt.Println(jsonKit.MarshalToString(src, jsonKit.WithApi(jsoniter.ConfigCompatibleWithStandardLibrary)))
	// {"b":false,"tmp":{"Id":666}} <nil>
	fmt.Println(jsonKit.MarshalToString(dest, jsonKit.WithApi(jsoniter.ConfigCompatibleWithStandardLibrary)))

	src["b"] = true
	b.Id = 777

	// {"b":true,"tmp":{"Id":777}} <nil>
	fmt.Println(jsonKit.MarshalToString(src, jsonKit.WithApi(jsoniter.ConfigCompatibleWithStandardLibrary)))
	// {"b":false,"tmp":{"Id":777}} <nil>
	fmt.Println(jsonKit.MarshalToString(dest, jsonKit.WithApi(jsoniter.ConfigCompatibleWithStandardLibrary)))
*/
func Copy(dest, src interface{}) error {
	return copier.CopyWithOption(dest, src, copier.Option{
		IgnoreEmpty: false,
		DeepCopy:    false,
		Converters:  nil,
	})
}

// DeepCopy 深拷贝（通过lancet）.
/*
PS: 不支持未导出的字段.

@param src	可以为nil（此时将返回nil）
*/
func DeepCopy[T any](src T) T {
	return convertor.DeepClone(src)
}

// DeepCopy1 深拷贝（通过GoFrame中的 gutil）.
/*
PS:
(1) 不支持未导出的字段 unable to copy unexported fields in a struct (lowercase field names)
(2) 不使用 github.com/mohae/deepcopy: 	虽然效果一样，但不推荐使用（star少; 最后更新时间2017）
(3) 不使用 github.com/jinzhu/copier: 	深拷贝有bug，详见"Golang.wps"

@param src	(1) 可以为nil（此时将返回: nil, nil）
			(2) 必须是: 结构体实例的指针 || map实例 || slice实例
*/
func DeepCopy1[T any](src T) (dest T, err error) {
	obj := gutil.Copy(src)
	var ok bool
	if dest, ok = obj.(T); ok {
		return
	}
	err = errorKit.Newf("Fail to deep copy because types of src(%T) and dest(%T) are different.", src, dest)
	return
}
