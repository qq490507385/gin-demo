package util

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// HmacSha256 hmac sha256加密
// 用于生成签名
func HmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// SaltSecret 随机数
func SaltSecret(len int) (string, error) {
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// SaltSecret16 16位随机数
func SaltSecret16() (string, error) {
	return SaltSecret(16)
}

// GenerateToken 生成Token
func GenerateToken(data string) (string, error) {
	t := time.Now().UnixNano()
	ts := strconv.FormatInt(t, 10)
	rn, err := SaltSecret16()
	if err != nil {
		return "", err
	}
	return HmacSha256(data, rn+ts), nil
}

// EncryptPassword 密码加密
func EncryptPassword(password string) (string, error) {
	encode, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encode), nil
}

// VerifyPassword 验证密码
func VerifyPassword(password string, hashed string) (isTrue bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return false, errors.New("Wrong password")
	}
	return true, nil
}
