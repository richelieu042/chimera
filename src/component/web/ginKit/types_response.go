package ginKit

type (
	ResponsePackage struct {
		// StatusCode http状态码（默认200，即 http.StatusOK）
		StatusCode int

		// Object json对象（不为nil的情况下，序列化为json并响应给前端）
		Object interface{}

		// Text 响应字符串（也可以是json）
		Text string

		// FilePath 响应文件的路径
		FilePath string
		// FileContent 响应文件的内容
		FileContent []byte
		// ContentType 响应文件内容的类型（需要与 FileContent 搭配使用）
		ContentType string
		// FileName 文件名
		/*
			PS: 使用 FilePath 的情况下， FileName 是可选的，为空的话会自己从 FilePath 中截取
		*/
		FileName string

		Error error
	}
)
