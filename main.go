package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Raj-deepp/RSS_Aggregator/internal/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080" // Default port
		log.Println("PORT not found, using default port 8080")
	}

	db_URL := os.Getenv("DB_URL")
	if db_URL == "" {
		log.Fatal("DB_URL not found")
	}

	conn, err := sql.Open("postgres", db_URL)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "OPTIONS", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	vR := router.Group("/v1")
	{
		vR.GET("/health", handlerReadiness)
		vR.GET("/err", handlerErr)

		vR.POST("/users", apiCfg.handlerUsersCreate)
		vR.GET("/users", apiCfg.middlewareAuth(apiCfg.handlerUsersGet))
	}

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portStr,
	}

	log.Printf("Server starting on port %v", portStr)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
