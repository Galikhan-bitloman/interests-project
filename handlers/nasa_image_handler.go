package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"refactor/model"
)

func (h *Handler) GetNasaImageHandler(ctx *fiber.Ctx) error {
	data, err := h.services.Nasa.GetNasaImage()

	if err != nil {
		log.Fatal(err)
	}

	res := model.DataResponse{
		Data: data,
	}

	log.Println(res)

	return ctx.JSON(&res)
}
