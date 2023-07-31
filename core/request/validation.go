package request

import (
	"admin/util"
	"github.com/go-playground/validator/v10"
	"os"
	"strings"
)

type validation struct {
}

var Validate = &validation{}

// FileIsExists 校验文件是否存在
func (v *validation) FileIsExists(fl validator.FieldLevel) bool {
	val, ok := fl.Field().Interface().(string)
	if ok && val != "" {
		_, err := os.Stat(val)
		if err != nil {
			if os.IsExist(err) {
				return true
			}
			return false
		}
	}
	return true
}

// Password 校验密码 6~32位，支持大小写字母、数字、~!@#$%&*，需包含2种类型以上
func (v *validation) Password(fl validator.FieldLevel) bool {
	val, ok := fl.Field().Interface().(string)
	if ok && val != "" {
		if len(val) < 6 && len(val) > 32 {
			return false
		}
		count := 0
		if strings.ContainsAny(val, util.StrUtil.Numbers) {
			count++
		}
		if strings.ContainsAny(val, util.StrUtil.Letters) {
			count++
		}
		if strings.ContainsAny(val, strings.ToUpper(util.StrUtil.Letters)) {
			count++
		}
		if strings.ContainsAny(val, util.StrUtil.Chars) {
			count++
		}
		return count >= 2
	}
	return true
}
