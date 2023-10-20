package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rahmatrdn/go-skeleton/internal/http/handler"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             200 * time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn,       // Log level
			IgnoreRecordNotFoundError: true,              // Ignore ErrRecordNotFound error for logger
		},
	)
	dsn := "root:@tcp(127.0.0.1:3306)/go_perpustakaan"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		AppName: "GO Perpustakaan v0.0.1",
	})

	apiV1 := app.Group("/v1")
	handler.NewUserHandler().Register(apiV1)

	log.Fatal(app.Listen(":3000"))
}
