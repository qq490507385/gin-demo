package exceptions

var (
	//OK  无异常
	OK = &Exception{Code: 0, Message: "Success"}
	//ServerError 系统错误  eg: 默认返回
	ServerError = &Exception{Code: 10001, Message: "Server error"}
	//ErrValidation  参数验证失败
	ErrValidation = &Exception{Code: 10002, Message: "Validation failed"}
	//ErrToken  签署JSON web令牌时出错
	ErrToken = &Exception{Code: 10003, Message: "Error occurred while signing the JSON web token."}
	//ErrAlExist 数据已存在
	ErrAlExist = &Exception{Code: 10004, Message: "Data already exists."}
	//ErrBind 参数绑定到结构体错误
	ErrBind = &Exception{Code: 10005, Message: "Error occurred while binding the request body to the struct."}
	// ErrDatabase 数据库操作失败
	ErrDatabase = &Exception{Code: 10006, Message: "Database operation failed."}

	//ErrPassword 帐号或密码错误
	ErrPassword = &Exception{Code: 20001, Message: "Wrong account number or password."}
)
