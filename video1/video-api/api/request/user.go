package request

type LoginUser struct {
	Mobile   string `form:"mobile" binding:"required" `   // 手机号
	Password string `form:"password" binding:"required" ` // 密码
}
