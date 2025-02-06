package handlers

import (
	"kurs/intrenal/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) InitRoute(app *fiber.App) *fiber.Handler {

	return nil
}
