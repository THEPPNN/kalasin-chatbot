package repository

import (
	"context"
	"kalasin-chatbot/config"
	"log"
)

func Save(sessionID string, role string, message string) {
	_, err := config.DB.Exec(
		context.Background(),
		"INSERT INTO chats(session_id, role, message) VALUES($1,$2,$3)",
		sessionID,
		role,
		message,
	)

	if err != nil {
		log.Println("DB INSERT ERROR:", err)
		return
	}

	log.Println("Saved chat →", sessionID, role, message)
}