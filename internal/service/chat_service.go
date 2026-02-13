package service

import (
    "context"
    "time"
    "kalasin-chatbot/config"
	"kalasin-chatbot/internal/ai"
    "kalasin-chatbot/internal/repository"
    "log"
)

func callLLM(ctx context.Context, messages []ai.Message) (string, error) {
	log.Println("LLM INPUT:", messages)

	// mock delay simulate AI
	select {
	case <-time.After(2 * time.Second):
		return "AI response mock", nil

	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func ChatAI(msg string) (string, error) {

	sessionID := "anonymous"
	repository.Save(sessionID, "user", msg)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	key := sessionID + ":" + msg

	val, err := config.RDB.Get(ctx, key).Result()

	if err == nil {
		log.Println("CACHE HIT")
		repository.Save(sessionID, "assistant", val)
		return val, nil
	}

	if err != redis.Nil {
		log.Println("Redis error:", err)
	}

	log.Println("CACHE MISS")

	history, err := repository.LoadRecent(sessionID, 5)
	if err != nil {
		log.Println("Memory load error:", err)
	}

	messages := ai.BuildPrompt(
		ai.SystemPrompt(),
		history,
		msg,
		1200,
	)

	reply := callLLM(messages)

	repository.Save(sessionID, "assistant", reply)

	if err := config.RDB.Set(ctx, key, reply, 10*time.Minute).Err(); err != nil {
		log.Println("Redis set error:", err)
	}

	return reply, nil
}