package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Routes(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"modules/home/html/home", gin.H{
			"title":  "home page",
			"app name": viper.Get("App.Name"),
		})
	})

	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK,"modules/home/html/about", gin.H{
			"title":  "About page",
			"app name": viper.Get("App.Name"),
		})
	})

}
