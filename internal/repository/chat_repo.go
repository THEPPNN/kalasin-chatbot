package repository

import (
	"context"
	"kalasin-chatbot/config"
	"log"
)

// requestID := uuid.NewString()
// log.Println("REQ:", requestID, "incoming:", msg)

func Save(sessionID, role, msg string) error {
	_, err := config.DB.Exec(
		context.Background(),
		"INSERT INTO chats(session_id, role, message) VALUES($1,$2,$3)",
		sessionID, role, msg,
	)
	return err
}

func LoadRecent(sessionID string, limit int) ([]map[string]string, error) {
	rows, err := config.DB.Query(
		context.Background(),
		"SELECT role, message FROM chats WHERE session_id=$1 ORDER BY created_at DESC LIMIT $2",
		sessionID,
		limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []map[string]string

	for rows.Next() {
		var role, message string
		if err := rows.Scan(&role, &message); err != nil {
			return nil, err
		}

		history = append(history, map[string]string{
			"role": role,
			"content": message,
		})
	}

	return history, nil
}