package user

// SignupResponse ...注册成功后返回数据绑定到此struct
type SignupResponse struct {
	Identifier string `json:"identifier"`
	UserID     string `json:"user_id"`
}

// SigninResponse ...登录成功后返回数据绑定到此struct
type SigninResponse struct {
	Token string `json:"token"`
}
