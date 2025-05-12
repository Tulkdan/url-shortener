package internal

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	var result []byte

	for range length {
		index := seededRand.Intn(len(charset))
		result = append(result, charset[index])
	}

	return string(result)
}
