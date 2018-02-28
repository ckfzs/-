package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "golang.org/x/sync/errgroup"
    "./app/test"
    "./app/auth"
    "./app/qrcode"
    "./common/GinHTMLRender"
    "net/http"
    _"net/http/pprof"
)

var (
    g errgroup.Group
)

func baseSetting(engine *gin.Engine) {
    engine.Static("/static", "./static")
    
    htmlRender := GinHTMLRender.New()
    htmlRender.Debug = gin.IsDebugging()
    htmlRender.TemplatesDir = "templates/"
    htmlRender.Ext = ".tmpl"
    htmlRender.Layout = "layouts/default"
    engine.HTMLRender = htmlRender.Create()
}

func main() {
    engine := gin.Default()
    baseSetting(engine)
    test.RegisterBlueprint(engine)
    auth.RegisterBlueprint(engine)
    qrcode.RegisterBlueprint(engine)
    g.Go(func() error {
        return engine.Run(":50000")
    })
    g.Go(func() error {
        return http.ListenAndServe(":50001", nil)
    })
    if err := g.Wait(); err != nil {
        fmt.Println(err)
    }
}