package main

import (
	"contact_log/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// TODO: Configura il middleware CORS per concentire le richieste da qualsiasi dominio, da sistemare per la connessione
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	r.Use(cors.New(config))

	r.GET("/contacts", controllers.GetContacts)
	r.GET("/contact/:id", controllers.GetContactByID)
	r.POST("/contact", controllers.CreateContact)
	r.PUT("/contact", controllers.UpdateContact)
	r.Run(":8080")
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Errore nel caricare il file .env: %s", err)
	}
}
