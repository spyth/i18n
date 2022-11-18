package main

import (
	"log"
	"net/http"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func main() {
	// new gin engine
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	localizer := ginI18n.Localize()

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, localizer.MustGetMessage("welcome", context.Query("lang")))
	})

	router.GET("/:name", func(context *gin.Context) {
		context.String(http.StatusOK, localizer.MustGetMessage(&i18n.LocalizeConfig{
			MessageID: "welcomeWithName",
			TemplateData: map[string]string{
				"name": context.Param("name"),
			},
		}, context.Query("lang")))
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
