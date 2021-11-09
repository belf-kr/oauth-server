package util

import (
	"crypto/sha256"
	"fmt"
)

// sha256으로 암호화 합니다.
func Sha256(v string) string {
	h := sha256.Sum256([]byte(v))
	return fmt.Sprintf("%x", h[:])
}
