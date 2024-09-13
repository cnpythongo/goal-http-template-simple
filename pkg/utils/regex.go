package utils

import (
	idvalidator "github.com/guanguans/id-validator"
	"github.com/jinzhu/copier"
	"regexp"
)

// CellphoneCheckRule 手机号验证规则
func CellphoneCheckRule() string {
	return "^1[3456789]{1}\\d{9}$"
}

// EmailCheckRule 邮箱验证规则
func EmailCheckRule() string {
	return "^[A-Z0-9._%+-]+@[A-Z0-9.-]+\\.[A-Z]{2,6}$"
}

// IsCellphone 校验手机号码规则
func IsCellphone(phone string) bool {
	regRuler := CellphoneCheckRule()
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(phone)
}

// IsEmail 校验邮箱规则
func IsEmail(email string) bool {
	regRuler := EmailCheckRule()
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(email)
}

// IsIdCardNumber 校验身份证号码规则，支持15位、18位及末位X，且会验证身份证号的合法性
func IsIdCardNumber(card string) bool {
	return idvalidator.IsStrictValid(card)
}

// FakeIdCartNumber 随机生成可通过校验的虚拟身份证号
func FakeIdCartNumber() string {
	return idvalidator.FakeId()
}

type IdCartInfo idvalidator.IdInfo

// GetIdCartInfo 严格模式获取身份证号信息
// []interface {}[
//
//	github.com/guanguans/id-validator.IdInfo{          // 身份证号信息
//	    AddressCode: int(500154)                           // 地址码
//	    Abandoned:   int(0)                                // 地址码是否废弃：1为废弃的，0为正在使用的
//	    Address:     string("重庆市市辖区开州区")             // 地址
//	    AddressTree: []string[                             // 省市区三级列表
//	        string("重庆市")                                    // 省
//	        string("市辖区")                                    // 市
//	        string("开州区")                                    // 区
//	    ]
//	    Birthday:      <1993-01-13 00:00:00 +0800 CST>     // 出生日期
//	    Constellation: string("摩羯座")                     // 星座
//	    ChineseZodiac: string("酉鸡")                       // 生肖
//	    Sex:           int(0)                              // 性别：1为男性，0为女性
//	    Length:        int(18)                             // 号码长度
//	    CheckBit:      string("6")                         // 校验码
//	}
//	<nil>                                              // 错误信息
//
// ]
func GetIdCartInfo(card string) (*IdCartInfo, error) {
	_info, err := idvalidator.GetInfo(card, true)
	if err != nil {
		return nil, err
	}
	var info IdCartInfo
	err = copier.Copy(info, _info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
