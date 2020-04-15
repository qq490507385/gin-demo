package request

import (
	"doc-system/app/models"
	"doc-system/util"

	"github.com/go-playground/validator/v10"
)

// UserCreateRequest 创建用户类型
type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Sex      int    `json:"set" validate:"oneof=0 1"` // 0 男  1 女
	Username string `json:"username" validate:"required,checkUsername"`
	Password string `json:"password" validate:"required"`
}

func validateMsgs() map[string]string {
	msg := util.ValidateMessages{}
	msg["Name.required"] = "name为必填字段"
	msg["Sex.oneof"] = "sex字段必须是0,1中一个"
	msg["Username.required"] = "username为必填字段"
	msg["Username.checkUsername"] = "username不能重复"
	msg["Password.required"] = "password为必填字段"
	return msg
}

// 自定义验证
func checkUsername(fl validator.FieldLevel) bool {
	userModel := models.UserModel{}
	val := fl.Field().String()
	if userModel.CheckUsernameExist(val) {
		return false
	}
	return true
}

// ValidateCreate 表单验证
func (u *UserCreateRequest) ValidateCreate() (msg util.ValidateMessages, e error) {
	validate := validator.New()
	validate.RegisterValidation("checkUsername", checkUsername)
	err := validate.Struct(u)
	if err != nil {
		validatorMsg := validateMsgs()
		messages := util.NewValidatorError(err, validatorMsg)
		return messages, err
	}
	return util.ValidateMessages{}, nil
}
