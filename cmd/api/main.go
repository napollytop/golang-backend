package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/napollytop/manga/internal/config"
	"github.com/napollytop/manga/internal/handler"
	"github.com/napollytop/manga/internal/models"
	"github.com/napollytop/manga/internal/routes"
)

func main(){
	db := config.ConnectDb()

	err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
	)
	if err != nil{
		log.Fatal("gagal migrasi dari struct models")
	}
	


	app := fiber.New()
	authHandler := handler.NewAuthHandler(db)
	productHandler := handler.NewProductHandler(db)
	
	routes.SetupAuthRoutes(app, authHandler)
	routes.GetProducts(app, productHandler)

	fmt.Println("server berjalan pada http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}