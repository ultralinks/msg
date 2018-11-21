package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func GetRandomNumber(length int) string {
	bytes := []byte("0123456789")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}

func GetRandomString(length int) string {
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}

func Md5(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	d := h.Sum(nil)

	return string(hex.EncodeToString(d))
}
