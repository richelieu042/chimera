package bcryptKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"golang.org/x/crypto/bcrypt"
)

// GenerateFromPassword 加密用户密码（生成一个带盐（salt）的哈希值）.
/*
@param costArgs 不传参将使用默认值(bcrypt.DefaultCost)

PS:
(1) 通常用于安全地存储密码;
(2) 传参相同，多次执行的结果不同.
*/
func GenerateFromPassword(password []byte, costArgs ...int) (hashedPassword []byte, err error) {
	/* 迭代次数，可以根据需要调整 */
	var cost int
	if costArgs != nil {
		cost = costArgs[0]
		if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
			err = errorKit.Newf("cost(%d) out of range([%d, %d])", cost, bcrypt.MinCost, bcrypt.MaxCost)
			return
		}
	} else {
		// 默认值
		cost = bcrypt.DefaultCost
	}

	hashedPassword, err = bcrypt.GenerateFromPassword(password, cost)
	return
}

// CompareHashAndPassword 验证密码是否正确（对比 已哈希过的密码 与 原始明文密码 是否匹配。）.
/*
@param hashedPassword 	GenerateFromPassword() 的返回值
@param plainPassword	密码明文
@return err == nil: 已哈希过的密码 与 原始明文密码 匹配
*/
func CompareHashAndPassword(hashedPassword, plainPassword []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, plainPassword)
}
