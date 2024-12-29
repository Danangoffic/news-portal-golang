package main

import (
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/http2"

	"news-portal/internal/database"
	"news-portal/internal/handlers"
	"news-portal/internal/middleware/AuthJWTNews"
	"news-portal/internal/models/article"
	"news-portal/internal/models/category"
	"news-portal/internal/repositories"
	"news-portal/internal/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectDB()
	_, err = database.GetDB()
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	// Migrasi database
	DB := database.DB
	err = DB.AutoMigrate(&category.Category{}, &article.Article{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	articleRepository := repositories.NewArticleRepository(DB)
	categoryRepository := repositories.NewCategoryRepository(DB)
	userRepository := repositories.NewUserRepository(DB)
	articleService := services.NewArticleService(articleRepository)
	categoryService := services.NewCategoryService(categoryRepository)
	userService := services.NewUserService(userRepository)
	articleHandler := handlers.NewArticleHandler(articleService)
	categoryHandler := handlers.NewCategoryHandler(*categoryService)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()
	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	c := jaegertracing.New(e, nil)
	defer c.Close()

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Article routes
	// articleRoutes := e.Group("/api/articles")
	e.POST("/api/login", userHandler.LoginHandler)

	e.GET("/api/articles", articleHandler.GetAllArticles)
	e.GET("/api/articles/:slug", articleHandler.GetArticleBySlug)
	e.GET("/api/articles/:id/detail", articleHandler.GetArticleById)
	e.GET("/api/articles/status/:status", articleHandler.GetArticlesByStatus)

	e.POST("/", articleHandler.CreateArticle, AuthJWTNews.AuthJWT)
	e.PUT("/api/articles/:id", articleHandler.UpdateArticle, AuthJWTNews.AuthJWT)
	e.DELETE("/api/articles/:id", articleHandler.DeleteArticle, AuthJWTNews.AuthJWT)

	// categoryRoutes := e.Group("/api/categories")
	e.GET("/api/categories/", categoryHandler.GetCategories)
	e.GET("/api/categories/:id", categoryHandler.GetCategory)

	e.POST("/api/categories/", categoryHandler.CreateCategory, AuthJWTNews.AuthJWT)
	e.PUT("/api/categories/:id", categoryHandler.UpdateCategory, AuthJWTNews.AuthJWT)
	e.DELETE("/api/categories/:id", categoryHandler.DeleteCategory, AuthJWTNews.AuthJWT)

	// Start server
	if err := e.StartH2CServer(":8080", s); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
