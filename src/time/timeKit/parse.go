package timeKit

import (
	"time"
)

// Parse 作用: string => time.Time
/*
PS: time.Parse使用 time.UTC 作为loc.

@param format 	时间格式
@param timeStr	要解析的时间字符串
*/
func Parse[F ~string](format F, timeStr string) (time.Time, error) {
	return time.Parse(string(format), timeStr)
}

// ParseInLocal
/*
PS: 如果 timeStr 中有时区，必须先使用 Parse 解析，再 .Local() 转换.

@param loc time.Local || time.UTC
*/
func ParseInLocal[F ~string](format F, timeStr string) (time.Time, error) {
	return time.ParseInLocation(string(format), timeStr, time.Local)
}

// ParseInLocation
/*
@param loc time.Local || time.UTC
*/
func ParseInLocation[F ~string](format F, timeStr string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(string(format), timeStr, loc)
}

// ParseDuration string => time.Duration
/*
@param str (1) 如果为 ""，将返回error
    	   (2) e.g. "300ms"、"-1.5h"、"2h45m"
*/
var ParseDuration func(str string) (time.Duration, error) = time.ParseDuration
