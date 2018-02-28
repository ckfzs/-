package auth

import (
    "fmt"
    "github.com/gin-gonic/gin"
    _"github.com/gin-gonic/gin/binding"
    "net/http"
)

func loginPage(c *gin.Context) {
    c.HTML(http.StatusOK, "login.tmpl", gin.H{
        "title": "test",
    })
}

func loginAPI(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")
    fmt.Printf("ckdebug: %s %s\n", username, password)
    if len(username) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "status": "failed", 
            "error": "Please enter your username",
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "status": "success",
            "username": username,
        })
    }
}

func RegisterBlueprint(engine *gin.Engine) {
    group := engine.Group("auth")
    {
        group.GET("/login", loginPage)
        group.POST("/login", loginAPI)
    }
    
}