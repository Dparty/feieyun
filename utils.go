package feieyun

import (
	"crypto/sha1"
	"encoding/hex"
)

func SHA1(str string) string {
	s := sha1.Sum([]byte(str))
	strsha1 := hex.EncodeToString(s[:])
	return strsha1
}
