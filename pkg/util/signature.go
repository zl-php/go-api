package util

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

// 根据签名串生成签名，用于简单的url参数签名
func GenParmSign(params map[string]string, key string) string {
	var keys []string
	var sorted []string
	for k, v := range params {
		if k != "sign" && v != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, k := range keys {
		sorted = append(sorted, fmt.Sprintf("%s=%s", k, params[k]))
	}
	str := strings.Join(sorted, "&")
	str += "&key=" + key
	return fmt.Sprintf("%X", md5.Sum([]byte(str)))
}
