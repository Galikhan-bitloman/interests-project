package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"refactor/model"
	"refactor/services"
)

type Handler struct {
	services  *services.Service
	customErr model.ErrorBody
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() (*fiber.App, error) {
	//	here's happening routing
	fmt.Println("here's happening routing")

	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")
	interests := v1.Group("/interests")
	interests.Get("/getNasaImage", h.GetNasaImageHandler)
	interests.Post("/calculate", h.SalaryCalculateHandler)

	return app, nil

}
