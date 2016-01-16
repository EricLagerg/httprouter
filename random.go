package httprouter

import (
	"math/rand"
	"time"

	"github.com/SermoDigital/helpers"
)

var src = rand.NewSource(time.Now().UnixNano())

const (
	n           = 32
	letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

// reqID returns a random hash + the timestamp in the format
// hash + ":" + time.UnixNano
func reqID() string {
	var b [1 + n + 64]byte // +1 colon

	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; remain-- {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
	}
	b[n] = ':'
	return string(b[:n+1+copy(b[n+1:], helpers.FormatUint(uint64(time.Now().UnixNano())))])
}
