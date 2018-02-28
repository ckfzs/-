package auth

type LoginForm struct {
    username string `form:"username" json:"username" binding:"required"`
    password string `form:"password" json:"password" binding:"required"`
}