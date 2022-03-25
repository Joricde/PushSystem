package util

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

func AddSalt(password string, salt int64) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	b := hash.Sum([]byte("1"))
	mid := hex.EncodeToString(b) + strconv.FormatInt(salt, 10)
	hash.Write([]byte(mid))
	result := hex.EncodeToString(b)
	return result
}
