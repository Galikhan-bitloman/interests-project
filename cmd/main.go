package main

import (
	"interests-project/pkg/handlers"
	"interests-project/pkg/repository"
	"log"

	"interests-project/pkg/services"
)

func main() {

	db, err := repository.ConnPg(repository.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		Username: "galikhan",
		Password: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Fatal("err in connection to pg", err)
	}

	repos := repository.NewRepository(db)
	service := services.NewService(repos)
	handler := handlers.NewHandler(service)
	routes, err := handler.InitRoutes()
	if err != nil {
		log.Fatal(err)
	}

	err = routes.Listen("0.0.0.0:3000")

	if err != nil {
		log.Fatal(err)
	}

}
