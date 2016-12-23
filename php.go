package php

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	md5Hash := md5Ctx.Sum(nil)
	md5Str := string(hex.EncodeToString(md5Hash))
	return md5Str
}
