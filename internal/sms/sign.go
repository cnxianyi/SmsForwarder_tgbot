package sms

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"time"
)

var Sign string

func GetSign() (int64, string) {
	now := time.Now()
	timestampMilli := now.UnixNano() / int64(time.Millisecond)

	str := fmt.Sprintf("%d\n%s", timestampMilli, Sign)

	// HmacSHA256
	h := hmac.New(sha256.New, []byte(Sign))
	h.Write([]byte(str))
	hashResult := h.Sum(nil)

	// Base64
	base64Result := base64.StdEncoding.EncodeToString(hashResult)

	// URL
	finalSign := url.QueryEscape(base64Result)

	return timestampMilli, finalSign
}
