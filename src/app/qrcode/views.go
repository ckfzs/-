package qrcode

import (
    _ "fmt"
    _ "reflect"
    "github.com/gin-gonic/gin"
    "net/http"
    "encoding/base64"
)


func qrcGeneratorPage(c *gin.Context) {
    c.HTML(http.StatusOK, "qrcode/generator", gin.H{
        "title": "QRCode",
    })
}

func qrcDecoderPage(c *gin.Context) {
    c.HTML(http.StatusOK, "qrcode/decoder", gin.H{
        "title": "QRCode",
    })
}

func qrcAPIDocPage(c *gin.Context) {
    c.HTML(http.StatusOK, "qrcode/api-doc", gin.H{
        "title": "QRCode",
    })
}

func qrcDecodeAPI(c *gin.Context) {
    file, _, err := c.Request.FormFile("qrcode")
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status": "failed",
            "error": err.Error(),
        })
        return
    }
    content, err := Decode(file)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status": "failed",
            "error": err.Error(),
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "status": "success",
            "content": content,
        })
    }
}

func qrcEncodeAPI(c *gin.Context) {
    content := c.PostForm("content")
    if len(content) == 0 {
        c.JSON(http.StatusOK, gin.H{
            "status": "failed",
            "error": "content can't be empty",
        })
        return
    }
    png_buf, err := Encode(content)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status": "failed",
            "error": err.Error(),
        })
    } else {
        pngBase64Str := base64.StdEncoding.EncodeToString(png_buf)
        c.JSON(http.StatusOK, gin.H{
            "status": "success",
            "type": "png",
            "base64": pngBase64Str,
        })
    }
}

func RegisterBlueprint(engine *gin.Engine) {
    group := engine.Group("qrcode")
    {
        group.GET("/generator", qrcGeneratorPage)
        group.GET("/decoder", qrcDecoderPage)
        group.GET("/api-doc", qrcAPIDocPage)
        group.POST("/decode", qrcDecodeAPI)
        group.POST("/encode", qrcEncodeAPI)
    }
    
}