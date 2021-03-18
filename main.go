package main

import (
	"fmt"
	"github.com/BigerWANG/my-gin/pkg/setting"
	"github.com/BigerWANG/my-gin/routers"
	"log"
	"net/http"
)

func main()  {
	// gin.Default 返回Gin的type Engine struct{...} 里边包含routerGroup
	// 相当于创建一个路由Handles,可以后期绑定各种路由规则和函数，中间件等
	//r:=gin.Default()
	// context 是 gin中的上下文，它允许代码在中间件之间传递变量，管理流，验证JSON请求，响应JSON
	// 请求等
	//r.GET("/test", func(context *gin.Context) {
	//
	//	context.JSON(200, gin.H{
	//		"wtf": "go die",
	//	})
	//})
	r := routers.InitRouter()
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: r,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1<<20,
	}
	err := s.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
}


