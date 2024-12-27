package main

import (
	"net/http"

	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"news-portal/internal/database"
	"news-portal/internal/handlers"
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
	// dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	database.ConnectDB()

	// database.DB = db
	database.DB.AutoMigrate(&category.Category{}, &article.Article{})

	articleRepository := repositories.NewArticleRepository(database.DB)
	articleService := services.NewArticleService(articleRepository)
	articleHandler := handlers.NewArticleHandler(articleService)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Article routes
	articles := e.Group("/articles")
	articles.GET("", articleHandler.GetArticle)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
