package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	store 		*session.Store
	AUTH_KEY 	string 	= "authenticated"
	USER_ID 	string 	= "user_id"
)

func Setup() {

	router := fiber.New()

	store = session.New(session.Config{
		CookieHTTPOnly: true,
		// CookieSecure: true, for https
		Expiration: time.Hour * 5,
	})

	router.Use(NewMiddleware(), cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins: "*",
		AllowHeaders: "Access-Control-Allow-Origin, Content-Type, Origin, Accept",
	}))

	router.Post("/auth/register", Register)
	router.Post("/auth/login", Login)
	router.Post("/auth/logout", Logout)
	router.Get("/auth/healthcheck", HealthCheck)

	router.Get("/user", GetUser)

	router.Listen(":5000")
}