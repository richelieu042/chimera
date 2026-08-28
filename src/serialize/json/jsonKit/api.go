package jsonKit

var (
	library string

	// defaultApi 默认的API
	defaultApi API = nil

	// stdApi 标准的API（会对map的keys排序）
	stdApi API = nil
)

type (
	API interface {
		Marshal(v any) ([]byte, error)

		MarshalIndent(v any, prefix, indent string) ([]byte, error)

		MarshalToString(v any) (string, error)

		Unmarshal(data []byte, v any) error

		UnmarshalFromString(str string, v any) error
	}
)

func GetLibrary() string {
	return library
}

func GetDefaultApi() API {
	return defaultApi
}

func GetStdApi() API {
	return stdApi
}

//func testAPI() {
//	api := defaultApi
//	m := map[string]interface{}{
//		"0": 3.1415926,
//		"1": 1,
//		"2": true,
//		"3": `~!@#$%^&*()_+{}|:"><?
//qwdqw强	\t\r\n无敌
//qwdqwd
//威尔法
//496465~·《》？L:"{}|,./l;'[]\/*-的确
//`,
//	}
//
//	jsonData, err := api.Marshal(m)
//	if err != nil {
//		logrus.WithError(err).Fatalf("[%s, JSON] Fail to marshal!!!", consts.UpperProjectName)
//	}
//	var m1 map[string]interface{}
//	if err := api.Unmarshal(jsonData, &m1); err != nil {
//		logrus.WithError(err).Fatalf("[%s, JSON] Fail to unmarshal!!!", consts.UpperProjectName)
//	}
//	if compareKit.EqualWithTypeAndValue(m, m1) {
//		diff := compareKit.Diff(m, m1)
//		logrus.WithField("diff", diff).Fatalf("[%s, JSON] m and m1 are different.", consts.UpperProjectName)
//	}
//}
