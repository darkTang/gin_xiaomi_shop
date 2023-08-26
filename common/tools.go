package common

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

// Md5 Md5加密
func Md5(str string) string {
	h := md5.New()
	_, _ = io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// UnixToTime 毫秒转化为时间
func UnixToTime(timestamp uint64) string {
	t := time.UnixMilli(int64(timestamp))
	return t.Format("2006-01-02 15:04:05")
}

func Atoi(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return v
}
