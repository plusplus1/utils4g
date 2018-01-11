package str

import (
	"crypto/md5"
	"fmt"
)

type Util struct{}

func (u *Util) Md5Bytes(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

func (u *Util) Md5String(data string) string {
	return u.Md5Bytes([]byte(data))
}
