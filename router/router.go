package router

import (
	"a21hc3NpZ25tZW50/handler"
	"a21hc3NpZ25tZW50/middleware"
	"a21hc3NpZ25tZW50/repository"
	"a21hc3NpZ25tZW50/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

func SetupRouter(client *resty.Client, token, baseUrl string, db *gorm.DB) *gin.Engine {
	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	aiRepo := repository.NewAiRepository(client, token, baseUrl)
	fileService := service.NewFileService(repository.FileRepository{})
	aiService := service.NewAiService(aiRepo)
	apiHandler := handler.NewHandler(aiService, fileService)

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// endpoint
	api := r.Group("/api/v1")

	app := api.Group("/app")
	app.Use(middleware.Auth())
	app.POST("/chat", apiHandler.ChatWithAI())
	app.POST("/analyzed", apiHandler.AnalyzeDataTable())

	auth := api.Group("/auth")
	auth.POST("/register", userHandler.Register)
	auth.POST("/login", userHandler.Login)
	auth.POST("/logout", middleware.Auth(), userHandler.Logout)

	return r
}
