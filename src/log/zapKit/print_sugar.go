package zapKit

func Debugf(template string, args ...interface{}) {
	getInnerSL().Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	getInnerSL().Debugw(msg, keysAndValues...)
}

func Debugln(args ...interface{}) {
	getInnerSL().Debugln(args...)
}

// Infof 格式化输出的信息日志，类似于 fmt.Printf ，可以使用格式化字符串.
func Infof(template string, args ...interface{}) {
	getInnerSL().Infof(template, args...)
}

// Infow 结构化输出的信息日志，使用键值对的方式输出，更加适合记录结构化数据.
/*
@param keysAndValues e.g. "key", "value", "flag", true
*/
func Infow(msg string, keysAndValues ...interface{}) {
	getInnerSL().Infow(msg, keysAndValues...)
}

// Infoln
/*
PS: Spaces are always added between arguments.（传参间会加上" "）
*/
func Infoln(args ...interface{}) {
	getInnerSL().Infoln(args...)
}

func Warnf(template string, args ...interface{}) {
	getInnerSL().Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	getInnerSL().Warnw(msg, keysAndValues...)
}

func Warnln(args ...interface{}) {
	getInnerSL().Warnln(args...)
}

func Errorf(template string, args ...interface{}) {
	getInnerSL().Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	getInnerSL().Errorw(msg, keysAndValues...)
}

func Errorln(args ...interface{}) {
	getInnerSL().Errorln(args...)
}

func DPanicf(template string, args ...interface{}) {
	getInnerSL().DPanicf(template, args...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	getInnerSL().DPanicw(msg, keysAndValues...)
}

func DPanicln(args ...interface{}) {
	getInnerSL().DPanicln(args...)
}

func Panicf(template string, args ...interface{}) {
	getInnerSL().Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	getInnerSL().Panicw(msg, keysAndValues...)
}

func Panicln(args ...interface{}) {
	getInnerSL().Panicln(args...)
}

func Fatalf(template string, args ...interface{}) {
	getInnerSL().Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	getInnerSL().Fatalw(msg, keysAndValues...)
}

func Fatalln(args ...interface{}) {
	getInnerSL().Fatalln(args...)
}
