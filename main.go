package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Initial RSS_Agg build")

	godotenv.Load()

	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT not found")
	}
	
	router := gin.New()
}
