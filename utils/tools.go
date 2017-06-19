package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

func Md5(str string) (result string) {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(cipherStr))
}
func GetTimeStr() string {
	return strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
}
func Int2Str(i int) string {
	return strconv.Itoa(i)
}
func StrJion(strs ...string) string {
	b := bytes.Buffer{}
	for _, v := range strs {
		b.WriteString(v)
	}
	return b.String()
}
