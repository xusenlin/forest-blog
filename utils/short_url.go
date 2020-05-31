package utils

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

const charset = "A0a12B3b4CDc56Ede7FGf8Hg9IhJKiLjkMNlmOPnQRopqrSstTuvUVwxWXyYzZ"

func generateCharset(url, hexMd5 string, len, sectionNum int, cb func(url, keyword string) bool) string {
	for i := 0; i < sectionNum; i++ {
		sectionHex := hexMd5[i*8 : 8+i*8]
		bits, _ := strconv.ParseUint(sectionHex, 16, 32)
		bits = bits & 0x3FFFFFFF
		keyword := ""
		for j := 0; j < len; j++ {
			idx := bits & 0x3D
			keyword = keyword + string(charset[idx])
			bits = bits >> 5
		}
		if cb(url, keyword) {
			return keyword
		}
	}
	return ""
}

// 起初生成6位的短码，当四组6位短码都重复时，再生成8位的短码，因此总共会有8个短码供你选择。

func GenerateShortUrl(url string, cb func(url, keyword string) bool) string {
	if url == "" || cb == nil {
		return ""
	}
	hexMd5 := fmt.Sprintf("%x", md5.Sum([]byte(url)))
	sections := len(hexMd5) / 8

	keyword := generateCharset(url, hexMd5, 6, sections, cb)
	if keyword == "" {
		return generateCharset(url, hexMd5, 8, sections, cb)
	}
	return keyword
}
