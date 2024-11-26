package validator

import (
	"reflect"
	"regexp"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 自定义验证器
var (
	// 手机号验证器
	mobileRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)
)

// Setup 设置验证器
func Setup() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器
		_ = v.RegisterValidation("mobile", validateMobile)
		_ = v.RegisterValidation("password", validatePassword)

		// 注册自定义标签名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

// validateMobile 验证手机号
func validateMobile(fl validator.FieldLevel) bool {
	return mobileRegex.MatchString(fl.Field().String())
}

// validatePassword 验证密码（至少包含数字和字母，长度8-20）
func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 || len(password) > 20 {
		return false
	}

	var hasLetter, hasNumber bool
	for _, r := range password {
		if unicode.IsLetter(r) {
			hasLetter = true
		} else if unicode.IsNumber(r) {
			hasNumber = true
		}
		if hasLetter && hasNumber {
			return true
		}
	}
	return false
}
