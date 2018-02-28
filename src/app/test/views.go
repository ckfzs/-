package test

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func testHandler(c *gin.Context) {
    c.String(http.StatusOK, "test api")
}

func RegisterBlueprint(engine *gin.Engine) {
    group := engine.Group("test")
    group.GET("/", testHandler)
}