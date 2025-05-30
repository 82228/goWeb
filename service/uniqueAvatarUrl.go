package service

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"path"
	"time"
)

// 生成唯一文件名
func GenerateUniqueFileName(originalName string) string {
	// 获取扩展名
	ext := path.Ext(originalName)

	// 生成随机字符串
	randomString := generateRandomString(8)

	// 生成时间戳
	timestamp := time.Now().Format("20060102150405")

	// 组合唯一文件名
	return fmt.Sprintf("%s_%s%s", timestamp, randomString, ext)
}

// 生成随机字符串
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[randInt(0, len(charset)-1)]
	}
	return string(b)
}

// 生成随机整数
func randInt(min int, max int) int {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		return min
	}
	return min + int(num.Int64())
}
