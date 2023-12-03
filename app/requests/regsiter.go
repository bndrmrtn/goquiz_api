package requests

import (
	"github.com/gofiber/fiber/v2"
)

type RegisterValidation struct {
	Username string `json:"username" validate:"required,min=3,max=15,alphanumunicode"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=10,max=25"`
}

func RegisterRequest(c *fiber.Ctx) error {
	r := &RegisterValidation{}
	return Validate(r, c)
}
