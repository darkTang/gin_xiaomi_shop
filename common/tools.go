package common

import (
	"crypto/md5"
	"fmt"
	"io"
)

// Md5 Md5加密
func Md5(str string) string {
	h := md5.New()
	_, _ = io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
