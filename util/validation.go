package util

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	// 校验中文化
	ch := zh.New()
	uni := ut.New(ch)
	trans, _ := uni.GetTranslator("zh")
	validate = validator.New()

	// 注册转换语言为默认语言
	zh_translations.RegisterDefaultTranslations(validate, trans)
}

// ValidateMessages 类型名称转换
type ValidateMessages map[string]string

// NewValidatorError 表单验证错误翻译
func NewValidatorError(err error, messages map[string]string) ValidateMessages {
	res := ValidateMessages{}
	errors := err.(validator.ValidationErrors)
	for _, err := range errors {
		lowField := strings.ToLower(err.Field())
		key := err.Field() + "." + err.Tag()
		if msg, ok := messages[key]; ok {
			res[lowField] = msg
		} else {
			res[lowField] = err.Translate(trans)
		}
	}
	return res
}
