package main

import (
    "github.com/gin-gonic/gin"
     "net/http"
)

func main(){
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.LoadHTMLGlob("pages/*.html")
    r.Static("/js","./pages/js")
    r.StaticFS("/css",http.Dir("./pages/css"))
    r.GET("/",func(c *gin.Context){
        c.HTML(200,"index.html",gin.H{})
    })
    r.Run()
    
}