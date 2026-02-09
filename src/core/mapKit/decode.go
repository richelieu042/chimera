package mapKit

import (
	"github.com/mitchellh/mapstructure"
	"github.com/richelieu042/chimera/v3/src/core/ptrKit"
)

// Decode 将 通用的map[string]interface{} 解码到对应的 Go结构体中 ，或者执行相反的操作。
/*
PS:
(1) 也可用于: map转map；
(2) 字段标签(tag)不同于json，详见: https://mp.weixin.qq.com/s/n_AXYv-p_ev-q23mhpDkOg .

@param input	structure（包括map）
@param output 	必须是（map或结构体的）指针
*/
func Decode(input interface{}, output interface{}) error {
	if err := ptrKit.AssertNotNilAndIsPointer(output); err != nil {
		return err
	}

	return mapstructure.Decode(input, output)
}
