package Tools

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func RandString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		index := seededRand.Intn(len(charset))
		b[i] = charset[index]
		if index == 0 || time.Now().Unix()%int64(index) == 0 {
			newIndex := seededRand.Intn(10)
			b[i] = charset[len(charset)-newIndex]
		}
	}
	return string(b)
}
