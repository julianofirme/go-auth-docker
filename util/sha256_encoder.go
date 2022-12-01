package util

import (
	"crypto/sha256"
	"fmt"
)

func SHA256Encoder(stringPass string) string {
	encoderPass := sha256.Sum256([]byte(stringPass))

	return fmt.Sprintf("%x", encoderPass)
}
