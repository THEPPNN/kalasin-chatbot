package handler

import (
    "kalasin-chatbot/internal/service"
    "github.com/gofiber/fiber/v2"
	"log"
)
// GET /history?session_id=xxx

func RegisterRoutes(app *fiber.App) {
    app.Post("/chat", Chat)
    app.Get("/history", History)
}
func Chat(c *fiber.Ctx) error {
	log.Println("📩 Incoming request to /chat")
    var body struct {
        Message string `json:"message"`
    }
	log.Println("BODY PARSER:", string(c.Body()))
    if err := c.BodyParser(&body); err != nil {
        return c.Status(400).JSON(err.Error())
    }

    reply, err := service.ChatAI(body.Message)

	if err != nil {
        return c.Status(500).JSON(err.Error())
    }

    return c.JSON(fiber.Map{
        "reply": reply,
    })
}
func History(c *fiber.Ctx) error {
    sessionID := c.Query("session_id")
    if sessionID == "" {
        return c.Status(400).JSON("session_id is required")
    }
    history := repository.GetHistory(sessionID)
    return c.JSON(history)
}