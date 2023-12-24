package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptPassword(password string) []byte {
	//decodeString, _ := base64.StdEncoding.DecodeString(password)
	//return decodeString
	ctx := md5.New()
	ctx.Write([]byte(password))
	return []byte(hex.EncodeToString(ctx.Sum(nil)))
}
