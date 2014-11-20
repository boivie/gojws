package gojws

import (
	"encoding/base64"
	"strings"
)

func B64encode(data []byte) string {
	return strings.Replace(
		base64.URLEncoding.EncodeToString(data),
		"=", "", -1)
}
