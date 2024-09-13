package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/pbkdf2"
	"math/big"
	"strings"
)

const NormalLetters string = "abcdefghjkmnopqrstuvwxyzABCDEFGHJKMNOPQRSTUVWXYZ123456789"

const NumberLetters string = "0123456789"

// _randomStr 生成随机字符串
func _randomStr(letters string, i int) string {
	var container string
	b := bytes.NewBufferString(letters)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for x := 0; x < i; x++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(NormalLetters[randomInt.Int64()])
	}
	return container
}

// RandomStr 生成随机字符串，默认长度8
func RandomStr(i ...int) string {
	num := 8
	if len(i) > 0 {
		num = i[0]
	}
	return _randomStr(NormalLetters, num)
}

// RandomNumberStr 生成随机数字字符串，默认长度6
func RandomNumberStr(i ...int) string {
	num := 6
	if len(i) > 0 {
		num = i[0]
	}
	return _randomStr(NumberLetters, num)
}

// Md5 获取字符串的md5值
func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

// encodePassword 生成密码加密串
func encodePassword(password, salt string) string {
	b := pbkdf2.Key([]byte(password), []byte(salt), 1000, 24, sha1.New)
	return hex.EncodeToString(b)
}

// GeneratePassword 生成密码串和盐
func GeneratePassword(password string) (string, string) {
	salt := RandomStr()
	hashPwd := encodePassword(password, salt)
	return hashPwd, salt
}

// VerifyPassword 校验密码
func VerifyPassword(password, hashPwd, salt string) bool {
	vp := encodePassword(password, salt)
	return vp == hashPwd
}

// UUID 生成不带中划线UUID
func UUID() string {
	uuidStr := uuid.New().String()
	return strings.ReplaceAll(uuidStr, "-", "")
}
