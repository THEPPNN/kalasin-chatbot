package main

import (
    "log"
    "kalasin-chatbot/config"
    "kalasin-chatbot/internal/handler"

    "github.com/gofiber/fiber/v2"
)

func main() {
    config.Init()
    config.InitDB()
    log.Println("DB Connected")
    config.InitRedis() // ← เพิ่มบรรทัดนี้
	log.Println("Redis Connected")
    app := fiber.New()
    handler.RegisterRoutes(app)

    log.Fatal(app.Listen(":3000"))
}