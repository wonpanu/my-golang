package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wonpanu/my-golang/pkg/entity"
	"github.com/wonpanu/my-golang/pkg/usecase"
)

type VaccineRoute struct {
	uc usecase.IVaccine
}

func (r VaccineRoute) GetAllVaccine(c *fiber.Ctx) error {
	reponse, err := r.uc.GetAll()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("fail to get all vaccines")
	}
	return c.Status(fiber.StatusOK).JSON(reponse)
}

func (r VaccineRoute) GetVaccineByID(c *fiber.Ctx) error {
	ID := c.Params("id")
	response, err := r.uc.GetByID(ID)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("fail to get a vaccine by id")
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (r VaccineRoute) CreateVaccine(c *fiber.Ctx) error {
	var body entity.Vaccine
	err := c.BodyParser(&body)
	if err != nil {
		log.Println(err, body)
		return c.Status(fiber.ErrBadRequest.Code).SendString("invalid payload")
	}
	reponse, err := r.uc.Create(body)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.ErrBadRequest.Code).SendString("fail to create a vaccine")
	}
	return c.Status(fiber.StatusOK).JSON(reponse)
}

func (r VaccineRoute) UpdateVaccine(c *fiber.Ctx) error {
	ID := c.Params("id")
	var body entity.Vaccine
	err := c.BodyParser(&body)
	if err != nil {
		log.Println(err, body)
		return c.Status(fiber.ErrBadRequest.Code).SendString("invalid payload")
	}
	response, err := r.uc.Update(ID, body)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.ErrBadRequest.Code).SendString("fail to update a vaccine")
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (r VaccineRoute) DeleteVaccine(c *fiber.Ctx) error {
	ID := c.Params("id")
	err := r.uc.Delete(ID)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.ErrBadRequest.Code).SendString("fail to delete a vaccine")
	}
	return c.Status(fiber.StatusOK).SendString("success")
}

func NewVaccineHandler(vaccineUsecase usecase.IVaccine) VaccineRoute {
	return VaccineRoute{
		uc: vaccineUsecase,
	}
}
