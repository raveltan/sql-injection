package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"user" gorm:"user"`
	Password string `json:"pass" gorm:"pass"`
}

type TokenResponse struct {
	Username string `json:"user"`
}

type TokenRequest struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&User{})
	db.Create(&User{Password: "aoligei", Username: "admin"})
	s := Server{Db: db}
	app := fiber.New()
	app.Post("/login", s.login)
	app.Listen(":8080")
}

type Server struct {
	Db *gorm.DB
}

func (s *Server) login(c *fiber.Ctx) error {
	var r TokenRequest
	if err := c.BodyParser(&r); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	var res User
	query := fmt.Sprintf("SELECT * FROM users WHERE username = (\"%v\") AND password = (\"%v\") limit 0,1", r.User, r.Pass)
	log.Println(query)
	s.Db.Raw(query).First(&res)
	if res.Username == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.JSON(TokenResponse{res.Username})
}
