package handler

import "github.com/gofiber/fiber/v2"

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) Register(app fiber.Router) {
	app.Get("/user", u.GetAllUser)
	app.Post("/user", u.CreateUser)
}

func (u *UserHandler) GetAllUser(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func (u *UserHandler) CreateUser(c *fiber.Ctx) error {
	// write create user data using gorm

	return c.SendString("Hello, World!")
}
