package utils

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestGenerateAndVerifyPassword(t *testing.T) {
	pwd := "123123"
	p, salt := GeneratePassword(pwd)
	fmt.Println(p)
	fmt.Println(salt)
	f := VerifyPassword(pwd, p, salt)
	assert.Equal(t, f, true)
}

// TestCaptchaNumberImage 测试用例：生成图形验证码
func TestCaptchaGenerateNumberImage(t *testing.T) {
	c := NewCaptcha(100, 40, 4)
	a, b, err := c.GenerateNumberImage()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(err)
	assert.NotEqual(t, a, "")
	assert.NotEqual(t, b, "")
	assert.Equal(t, err, nil)
}

func TestCaptchaGenerateAlphabetImage(t *testing.T) {
	c := NewCaptcha(100, 40, 4)
	a, b, err := c.GenerateLettersImage()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(err)
	assert.NotEqual(t, a, "")
	assert.NotEqual(t, b, "")
	assert.Equal(t, err, nil)
}

func TestCaptchaGenerateMathImage(t *testing.T) {
	c := NewCaptcha(100, 40, 4)
	a, b, err := c.GenerateMathImage()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(err)
	assert.NotEqual(t, a, "")
	assert.NotEqual(t, b, "")
	assert.Equal(t, err, nil)
}
