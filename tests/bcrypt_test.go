package tests

import (
	"doc-system/util"
	"testing"
)

func TestSalt16(t *testing.T) {
	salt, err := util.SaltSecret16()
	if err != nil {
		t.Log(err)
	}
	t.Log(salt)
}

func TestGenerateToken(t *testing.T) {
	var data = "hello"
	token, err := util.GenerateToken(data)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(token)
	}
}

func TestEncrptPassword(t *testing.T) {
	var password = "123456"
	pass, err := util.EncryptPassword(password)
	if err != nil {
		t.Log(err)
	}
	t.Log(pass)
}

func TestVrifyPassword(t *testing.T) {
	var hashed = "$2a$10$6LqPWp3AdEL9r2slNulrkeeiCkqCk5wkUL1QE7FcGRD.j2hmy0OjK"
	var password = "123456"
	if ok, _ := util.VerifyPassword(password, hashed); ok {
		t.Log("密码正确")
	} else {
		t.Log("密码错误")
	}

}
