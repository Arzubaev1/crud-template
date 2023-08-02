package api

import (
	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageI, logger logger.LoggerI) {

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization
	handler := handler.NewHandler(cfg, storage, logger)

	r.Use(customCORSMiddleware())

	v1 := r.Group("/v1")

	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetByIdCategory)
	r.GET("/category", handler.GetListCategory)
	r.PUT("/category/:id", handler.UpdateCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)

	r.POST("/product", handler.CreateProduct)
	r.GET("/product/:id", handler.GetByIdProduct)
	r.GET("/product", handler.GetListProduct)
	r.PUT("/product/:id", handler.UpdateProduct)
	r.PATCH("/product/:id", handler.PatchProduct)
	r.DELETE("/product/:id", handler.DeleteProduct)

	v1.Use(handler.AuthMiddleware())
	v1.POST("/market", handler.CreateMarket)
	v1.GET("/market/:id", handler.GetByIdMarket)
	v1.GET("/market", handler.GetListMarket)
	v1.PUT("/market/:id", handler.UpdateMarket)
	v1.PATCH("/market/:id", handler.PatchMarket)
	v1.DELETE("/market/:id", handler.DeleteMarket)

	r.POST("/user", handler.CreateUser)
	r.GET("/user/:id", handler.AuthMiddleware(), handler.GetByIdUser)
	r.GET("/user", handler.GetListUser)
	r.PUT("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Accesp-Encoding, Authorization, Cache-Control")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
