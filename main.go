package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"pintureria-lincoln/internal/handler"
	"pintureria-lincoln/internal/repository"
)

func main() {
	exec, _ := os.Executable()
	_ = godotenv.Load(filepath.Join(filepath.Dir(exec), ".env"))
	_ = godotenv.Load()

	db := connectDB()
	defer db.Close()

	productRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	productHandler := handler.NewProductHandler(productRepo)
	orderHandler := handler.NewOrderHandler(orderRepo, productRepo)

	r := gin.Default()
	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("web/templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	api := r.Group("/api")
	{
		api.GET("/products", productHandler.GetAll)
		api.GET("/products/:id", productHandler.GetByID)
		api.GET("/categories", productHandler.GetCategories)
		api.POST("/orders", orderHandler.Create)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("servidor corriendo en http://localhost:%s", port)
	r.Run(":" + port)
}

func connectDB() *sql.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = fmt.Sprintf(
			"postgres://%s@%s:%s/%s?sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("error al abrir DB: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("error al conectar DB: %v", err)
	}
	log.Println("conectado a PostgreSQL")
	return db
}
