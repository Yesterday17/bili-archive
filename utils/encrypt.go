package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(text string) string {
	sign := md5.New()
	sign.Write([]byte(text))
	return fmt.Sprintf("%x", sign.Sum(nil))
}
