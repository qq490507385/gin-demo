package users

// LoginRequest 登录类型
type LoginRequest struct {
}

// CreateUserRequest 新增用户类型
type CreateUserRequest struct {
	Name     string `json:"name"`
	Sex      int    `json:"sex"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// UdapteUserRequest 更新用户类型
type UdapteUserRequest struct {
	ID       string `json:"id"`
	Sex      int    `json:"sex"`
	Username string `json:"useranme"`
}

// UpdatePasswordRequest 更新用户密码类型
type UpdatePasswordRequest struct {
	UdapteUserRequest
	Password string `json:"password"`
}
