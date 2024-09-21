package utils

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type Captcha struct {
	Width  int
	Height int
	Length int

	store   base64Captcha.Store
	bgColor color.RGBA
	fonts   []string
}

func NewCaptcha(width, height, length int) *Captcha {
	return &Captcha{
		Width:  width,
		Height: height,
		Length: length,

		store:   base64Captcha.DefaultMemStore,
		bgColor: color.RGBA{R: 233, G: 233, B: 233, A: 80},
		fonts:   []string{"wqy-microhei.ttc"},
	}
}

func (c *Captcha) SetStore(store base64Captcha.Store) {
	c.store = store
}

func (c *Captcha) SetBgColor(rgba color.RGBA) {
	c.bgColor = rgba
}

func (c *Captcha) SetFonts(fonts []string) {
	c.fonts = fonts
}

// GenerateNumberImage 生成全数字验证码和图片
func (c *Captcha) GenerateNumberImage() (code string, b64s string, err error) {
	driver := &base64Captcha.DriverDigit{
		Height:   c.Height,
		Width:    c.Width,
		Length:   c.Length,
		MaxSkew:  0.2,
		DotCount: 50,
	}
	return c.generate(driver)
}

// GenerateLettersImage 生成数字与字母混合的验证码和图片
func (c *Captcha) GenerateLettersImage() (code string, b64s string, err error) {
	driverStr := base64Captcha.DriverString{
		Height:          c.Height,
		Width:           c.Width,
		NoiseCount:      30,
		ShowLineOptions: 0,
		Length:          c.Length,
		Source:          NormalLetters,
		BgColor:         &c.bgColor,
		Fonts:           c.fonts,
	}
	driver := driverStr.ConvertFonts()
	return c.generate(driver)
}

// GenerateMathImage 生成算术验证码和图片
func (c *Captcha) GenerateMathImage() (code string, b64 string, err error) {
	driverStr := base64Captcha.DriverMath{
		Height:          c.Height,
		Width:           c.Width,
		NoiseCount:      0,
		ShowLineOptions: 0,
		BgColor:         &c.bgColor,
		Fonts:           c.fonts,
	}
	driver := driverStr.ConvertFonts()
	return c.generate(driver)
}

func (c *Captcha) generate(driver base64Captcha.Driver) (code string, b64s string, err error) {
	nc := base64Captcha.NewCaptcha(driver, c.store)
	code, b64s, _, err = nc.Generate()
	if err != nil {
		return "", "", err
	}
	return
}

// Verify 校验应答的验证码,校验完成无论是否正确都会清理缓存
// param: id string "生成验证码时的ID"
// param: answer string "应答码"
func (c *Captcha) Verify(id, answer string) (match bool) {
	return c.store.Get(id, true) == answer
}
