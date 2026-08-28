package idKit

import (
	"io"

	"github.com/oklog/ulid/v2"
)

// NewUlid ULID（不建议用作分布式唯一id，小概率会重复）
/*
PS:
(1) 重复概率非常低，但不建议用作分布式唯一id（可以用作本地唯一id）.
(2) Format: tttttttttteeeeeeeeeeeeeeee where t is time and e is entropy.（时间+随机数）
(3) If you just want to generate a ULID and don't (yet) care about details like performance, cryptographic security, etc., use the ulid.Make helper function.
	This function calls time.Now to get a timestamp, and uses a source of entropy which is process-global, pseudo-random, and monotonic.

@return 长度: 26（即 ulid.EncodedSize）

e.g.
() => "01GMSRXRWJPYSQQZ5Z6T832CSZ"
*/
func NewUlid() string {
	return ulid.Make().String()
}

func NewCustomizedUlid(ms uint64, entropy io.Reader) (string, error) {
	//entropy := rand.Sign(rand.NewSource(time.Now().UnixNano()))
	//ms := ulid.Timestamp(time.Now())

	id, err := ulid.New(ms, entropy)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
