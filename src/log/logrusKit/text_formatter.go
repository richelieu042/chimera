package logrusKit

import (
	"runtime"

	"github.com/richelieu042/chimera/v3/src/core/strKit"
	"github.com/richelieu042/chimera/v3/src/time/timeKit"
	"github.com/sirupsen/logrus"
)

func NewDefaultTextFormatter() *logrus.TextFormatter {
	return NewTextFormatter("")
}

// NewTextFormatter
/*
PS: 外部在调用此方法后，建议调用: Logger.SetReportCaller(true)!!!

@param timestampFormat 可以为""（将采用默认值）

e.g. 日志输出
time=2023-03-23 16:46:23.398+08:00 level=info msg=[CHIMERA, PROCESS] pid: [8579]. func=PrintBasicDetails(logrusKit/basicDetails.go:17)
time=2023-03-23 16:46:23.398+08:00 level=info msg=[CHIMERA, OS] os: [darwin]. func=PrintBasicDetails(logrusKit/basicDetails.go:20)
time=2023-03-23 16:46:23.398+08:00 level=info msg=[CHIMERA, OS] arch: [arm64]. func=PrintBasicDetails(logrusKit/basicDetails.go:21)
*/
func NewTextFormatter(timestampFormat string) *logrus.TextFormatter {
	if strKit.IsEmpty(timestampFormat) {
		timestampFormat = string(timeKit.FormatRFC3339)
	}

	// 是否用""将字段的值包起来?
	quoteFlag := true
	// Richelieu: 此处暂时固定为 false ，否则输出会很难看
	colorFlag := false

	return &logrus.TextFormatter{
		/* 时间格式 */
		TimestampFormat: timestampFormat,
		/* 禁止显示时间 */
		DisableTimestamp: false,
		/* 显示完整时间 */
		FullTimestamp: true,

		/* 禁止颜色显示 */
		DisableColors: !colorFlag,
		ForceColors:   colorFlag,

		DisableQuote: !quoteFlag,
		ForceQuote:   quoteFlag,

		// Disables the truncation of the level text to 4 characters.
		DisableLevelTruncation: false,
		// PadLevelText Adds padding the level text so that all the levels output at the same length
		// PadLevelText is a superset of the DisableLevelTruncation option
		PadLevelText: false,

		QuoteEmptyFields: true,

		CallerPrettyfier: func(f *runtime.Frame) (funcName string, fileName string) {
			return GetLocation(f), ""

			//s := strings.Split(f.Function, ".")
			//funcName = s[len(s)-1]
			//
			//s1 := strKit.Split(f.File, "/")
			//length := len(s1)
			//if length >= 2 {
			//	fileName = fmt.Sprintf("%s/%s:%d", s1[length-2], s1[length-1], f.Line)
			//} else {
			//	fileName = fmt.Sprintf("%s:%d", f.File, f.Line)
			//}
			//
			//// 把 file属性 整合到 func属性 里
			//funcName = fmt.Sprintf("%s(%s)", funcName, fileName)
			//// 不输出 file属性
			//fileName = ""
			//
			//return funcName, fileName
		},
	}
}
