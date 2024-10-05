package utils

import (
	"encoding/json"
	"math"
	"os"
	"reflect"
)

// toolsUtil 常用工具集合类
type toolsUtil struct{}

var (
	ToolsUtil = &toolsUtil{}
)

// Contains 判断src是否包含elem元素
func (t *toolsUtil) Contains(src interface{}, elem interface{}) bool {
	srcArr := reflect.ValueOf(src)
	if srcArr.Kind() == reflect.Ptr {
		srcArr = srcArr.Elem()
	}
	if srcArr.Kind() == reflect.Slice {
		for i := 0; i < srcArr.Len(); i++ {
			if srcArr.Index(i).Interface() == elem {
				return true
			}
		}
	}
	return false
}

// Round float四舍五入
func (t *toolsUtil) Round(val float64, n int) float64 {
	base := math.Pow(10, float64(n))
	return math.Round(base*val) / base
}

// JsonToObj JSON转Obj
func (t *toolsUtil) JsonToObj(jsonStr string, toVal interface{}) (err error) {
	return json.Unmarshal([]byte(jsonStr), &toVal)
}

// ObjToJson Obj转JSON
func (t *toolsUtil) ObjToJson(data interface{}) (res string, err error) {
	b, err := json.Marshal(data)
	if err != nil {
		return res, err
	}
	res = string(b)
	return res, nil
}

// IsFileExist 判断文件或目录是否存在
func (t *toolsUtil) IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
