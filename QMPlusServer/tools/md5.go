package tools

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5V ...
func MD5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
