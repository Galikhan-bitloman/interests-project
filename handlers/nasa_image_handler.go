package handlers

import (
	"log"
	"refactor/model"
	"refactor/sqlscripts"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetNasaImageHandler(ctx *fiber.Ctx) error {
	data, err := h.services.Nasa.GetNasaImageService()

	if err != nil {
		log.Fatal(err)
	}

	res := model.DataResponse{
		Data: data,
	}

	log.Println(res)

	return ctx.JSON(&res)
}

func (h *Handler) GetAllNasaImage(ctx *fiber.Ctx) error {
	data, err := h.services.GetAllNasaImageService(sqlscripts.GetAllNasaImage)

	if err != nil {
		return err
	}

	res := model.DataResponse{
		Data: data,
	}

	log.Println(res)

	return ctx.JSON(&res)

}
