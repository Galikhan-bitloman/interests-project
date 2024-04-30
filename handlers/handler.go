package handlers

import (
	"fmt"
	"refactor/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() (*fiber.App, error) {
	//	here's happening routing
	fmt.Println("here's happening routing")

	app := fiber.New()

	app.Use(cors.New())
	app.Use(csrf.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")
	interests := v1.Group("/interests")
	interests.Get("/getNasaImage", h.GetNasaImageHandler)
	interests.Post("/calculate", h.SalaryCalculateHandler)
	interests.Get("/all/nasaImage", h.GetAllNasaImage)

	return app, nil

}
