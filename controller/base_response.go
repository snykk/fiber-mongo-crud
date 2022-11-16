package controller

import "github.com/gofiber/fiber/v2"

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(ctx *fiber.Ctx, message string, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(BaseResponse{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func NewErrorResponse(ctx *fiber.Ctx, code int, message string) error {
	return ctx.Status(code).JSON(BaseResponse{
		Status:  false,
		Message: message,
	})
}
