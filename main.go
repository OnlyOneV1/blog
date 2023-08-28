package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/thinkerou/favicon"
)

// 中间件
func myHandler() gin.HandlerFunc { // handlerfunc表示返回自己的中间件
	return func(ctx *gin.Context) {
		ctx.Set("usersesion", "userid-1")
		ctx.Next()
		ctx.Abort()
	}

}

func main() {
	ginServer := gin.Default()
	// ginServer.Use(favicon.New("./favicon.ico"))

	ginServer.LoadHTMLGlob("templates/*")

	ginServer.GET("/index", func(ctx *gin.Context) {
		// gin.H{}是一个MAP字典
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "hello houtai!",
		})
	})

	// 添加中间件
	ginServer.GET("/user/info", myHandler(), func(ctx *gin.Context) {
		usersesion := ctx.MustGet("usersesion").(string)
		log.Println("==========>", usersesion)

		userid := ctx.Query("userid")
		username := ctx.Query("username")
		ctx.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})
	// user/info?userid=2&username=3
	// ginServer.GET("/user/info", func(ctx *gin.Context) {
	// 	userid := ctx.Query("userid")
	// 	username := ctx.Query("username")
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"userid":   userid,
	// 		"username": username,
	// 	})
	// })
	// // user/info/12/11
	// ginServer.GET("/user/info/:userid/:username", func(ctx *gin.Context) {
	// 	userid := ctx.Param("userid")
	// 	username := ctx.Param("username")
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"userid":   userid,
	// 		"username": username,
	// 	})
	// })
	// 前端给后端传送数据
	ginServer.POST("/json", func(ctx *gin.Context) {
		data, _ := ctx.GetRawData() // GetRawData 是要接受的数据 赋予data 是byte[]数据
		var m map[string]interface{}
		_ = json.Unmarshal(data, &m) // 将[]byte包装为json数据
		ctx.JSON(http.StatusOK, m)
	})
	//ginServer.Run(":8082")

	// 表单
	ginServer.POST("/user/add", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		ctx.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
		})
	})

	// 重定向Redirect
	ginServer.GET("/test", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	ginServer.Run(":8082")

	// 路由组
	// userGroup := ginServer.Group("/user")
	// {
	// 	userGroup.GET("/add")
	// 	userGroup.POST("/login")
	// 	userGroup.POST("/logout")
	// }
	// orderGroup := ginServer.Group("/order"){
	// 	orderGroup.GET("/add")
	// 	orderGroup.DELETE("/delete")
	// }

}
