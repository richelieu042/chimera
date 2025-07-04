package carbonKit

import (
	"github.com/dromara/carbon/v2"
)

var (
	// CreateFromStdTime time.Time => carbon.Carbon
	CreateFromStdTime = carbon.CreateFromStdTime

	// CreateFromDate 从给定的年、月、日创建 Carbon 实例
	CreateFromDate      = carbon.CreateFromDate
	CreateFromDateMilli = carbon.CreateFromDateMilli
	CreateFromDateMicro = carbon.CreateFromDateMicro
	CreateFromDateNano  = carbon.CreateFromDateNano

	// CreateFromDateTime 从给定的年、月、日、时、分、秒创建 Carbon 实例
	CreateFromDateTime      = carbon.CreateFromDateTime
	CreateFromDateTimeMilli = carbon.CreateFromDateTimeMilli
	CreateFromDateTimeMicro = carbon.CreateFromDateTimeMicro
	CreateFromDateTimeNano  = carbon.CreateFromDateTimeNano

	// CreateFromTimestamp 从给定的秒级时间戳创建 Carbon 实例
	CreateFromTimestamp      = carbon.CreateFromTimestamp
	CreateFromTimestampMilli = carbon.CreateFromTimestampMilli
	CreateFromTimestampMicro = carbon.CreateFromTimestampMicro
	CreateFromTimestampNano  = carbon.CreateFromTimestampNano

	// CreateFromTime 从给定的时、分、秒创建 Carbon 实例
	CreateFromTime      = carbon.CreateFromTime
	CreateFromTimeMilli = carbon.CreateFromTimeMilli
	CreateFromTimeMicro = carbon.CreateFromTimeMicro
	CreateFromTimeNano  = carbon.CreateFromTimeNano
)
