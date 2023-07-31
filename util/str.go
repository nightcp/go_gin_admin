package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type strUtil struct {
	Numbers string // 数字
	Letters string // 字母
	Chars   string
}

var (
	StrUtil = &strUtil{
		Numbers: "0123456789",
		Letters: "abcdefghijklmnopqrstuvwxyz",
		Chars:   "~!@#$%&*",
	}
	allRandomStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// MakeRandomStr 生成指定位数的随机字符串
func (su strUtil) MakeRandomStr(length int, str ...string) string {
	base := allRandomStr
	if len(str) > 0 {
		base = str[0]
	}
	byteList := make([]byte, length)
	for i := 0; i < length; i++ {
		byteList[i] = base[rand.Intn(len(base))]
	}
	return string(byteList)
}

// MakeMD5 生成MD5字符串
func (su strUtil) MakeMD5(str string) string {
	sum := md5.Sum([]byte(str))
	return hex.EncodeToString(sum[:])
}

// MakePassword 生成密码
func (su strUtil) MakePassword(str string, salt string) string {
	return su.MakeMD5(su.MakeMD5(str+salt) + salt)
}

// CheckPassword 验证密码
func (su strUtil) CheckPassword(str string, salt string, password string) bool {
	return su.MakePassword(str, salt) == password
}
