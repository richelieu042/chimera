package strKit

import "github.com/samber/lo"

var (
	// Ellipsis 将字符串截断为指定长度，并在截断后附加省略号.
	/*
		e.g.
		("Richelieu", 0) => "..."
		("Richelieu", 1) => "..."
		("Richelieu", 2) => "..."
		("Richelieu", 3) => "..."
		("Richelieu", 4) => "R..."
		("Richelieu", math.MaxInt) => "Richelieu"
	*/
	Ellipsis func(str string, length int) string = lo.Ellipsis
)
