package main

import (
	"os"

	"github.com/richelieu042/chimera/v3/src/file/fileKit"
	"github.com/richelieu042/chimera/v3/src/log/console"
	"github.com/richelieu042/chimera/v3/src/time/timeKit"
	"go.uber.org/zap"
)

func main() {
	path := "/Users/richelieu/Downloads"

	info, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	console.Debug("ModTime", zap.String("time", timeKit.Format(info.ModTime(), timeKit.FormatCommon)))

	t, ok := fileKit.GetBirthTime(info)
	if !ok {
		panic("GetBirthTime failed")
	}
	console.Debug("BirthTime", zap.String("time", timeKit.Format(t, timeKit.FormatCommon)))

}
