package common

import (
	"bytes"
	"encoding/json"
)

// FormatJSON 格式化为json
func FormatJSON(b []byte) string {
	// 转码格式化
	var out bytes.Buffer
	_ = json.Indent(&out, b, "", "\t")
	return out.String()
}
