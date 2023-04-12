package router

import (
	"net/http"

	"github.com/LinkinStars/dc/internal/base/config"
	"github.com/LinkinStars/dc/internal/base/middleware"
	"github.com/LinkinStars/dc/internal/controller"
	"github.com/LinkinStars/dc/static"
	"github.com/LinkinStars/go-scaffold/logger"
	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(middleware.CORS())
	r.StaticFS("/card", http.FS(static.Build))
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/card/page")
	})
	r.GET("/favicon.ico", func(ctx *gin.Context) {
		ctx.Header("content-type", "image/x-icon")
		file, err := static.Build.ReadFile("favicon.ico")
		if err != nil {
			logger.Error(err)
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.String(http.StatusOK, string(file))
	})

	r.NoRoute(func(ctx *gin.Context) {
		ctx.Header("content-type", "text/html;charset=utf-8")
		file, err := static.Build.ReadFile("index.html")
		if err != nil {
			logger.Error(err)
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.String(http.StatusOK, string(file))
	})

	v1 := r.Group("/api/v1")
	v1.POST("/card/login", controller.Login)

	authV1 := r.Group("/api/v1")
	authV1.Use(middleware.Auth(false))
	authV1.GET("/card/site", controller.GetSiteInfo)
	authV1.GET("/card", controller.GetCardDetail)
	authV1.GET("/cards", controller.GetCardsPage)
	authV1.GET("/cards/stat", controller.GetCardsStat)

	mustAuthV1 := r.Group("/api/v1")
	mustAuthV1.Use(middleware.Auth(true))

	mustAuthV1.POST("/card", controller.PostCard)
	mustAuthV1.PUT("/card", controller.UpdateCard)
	mustAuthV1.DELETE("/card", controller.DeleteCard)

	logger.Debugf("daily card run at %s", config.GlobalConf.WebPort)
	err := r.Run(config.GlobalConf.WebPort)
	if err != nil {
		panic(err)
	}
}
