package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"refactor/model"
)

func (h *Handler) SalaryCalculateHandler(ctx *fiber.Ctx) error {

	s := new(model.SalaryRequest)
	err := ctx.BodyParser(s)

	if err != nil {
		return err
	}

	data, err := h.services.SalaryCalculation.CalculateSalaryService(*s)

	resp := *data

	res := model.DataSalaryResponse{
		Data: &resp,
	}
	log.Println(res)

	return ctx.JSON(res)
}
