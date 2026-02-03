package zapKit

import (
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

// 自定义编码器来给日志消息添加前缀
type prefixEncoder struct {
	zapcore.Encoder

	prefix string
}

func (pe *prefixEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	// 给msg字段加上前缀
	entry.Message = pe.prefix + entry.Message
	return pe.Encoder.EncodeEntry(entry, fields)
}

// Clone !!!: 必须要重写(Overriding)此方法，否则会有问题，e.g.后续调用 zapcore.ioCore.With().
func (pe *prefixEncoder) Clone() zapcore.Encoder {
	return &prefixEncoder{
		Encoder: pe.Encoder.Clone(),
		prefix:  pe.prefix,
	}
}

// attachPrefixToEncoder 会给msg字段加上前缀.
/*
@param encoder 不能为nil
@return 可能是传参encoder
*/
func attachPrefixToEncoder(encoder zapcore.Encoder, prefix string) zapcore.Encoder {
	if prefix == "" {
		return encoder
	}

	if pe, ok := encoder.(*prefixEncoder); ok {
		pe.prefix = prefix
		return pe
	}
	return &prefixEncoder{
		Encoder: encoder,
		prefix:  prefix,
	}
}
