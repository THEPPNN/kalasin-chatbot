package service

import (
    "context"
    "time"
    "kalasin-chatbot/config"
    "kalasin-chatbot/internal/repository"
    "log"
)

func ChatAI(msg string) (string, error) {
	log.Println("🔥 HIT CHAT AI")
	sessionID := "anonymous"
    // save user message immediately (log every request)
    repository.Save(sessionID, "user", msg)
    // 1️⃣ check cache ก่อน
    log.Println("STEP REDIS GET")

    if config.RDB == nil {
        log.Println("Redis client is NIL")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    log.Println("PING REDIS...")
    if pingErr := config.RDB.Ping(ctx).Err(); pingErr != nil {
        log.Println("Redis ping failed:", pingErr)
    } else {
        log.Println("Redis ping OK")
    }

    log.Println("CALL REDIS GET NOW")
    val, err := config.RDB.Get(ctx, msg).Result()
    log.Println("RETURN REDIS GET")

    if err != nil {
        log.Println("Redis error:", err)
    } else {
        log.Println("Redis value:", val)
    }
	if err == nil && val != "" {
      log.Println("CACHE HIT → still logging to DB")
      repository.Save(sessionID, "assistant", val)
      return val, nil
	}
    log.Println("STEP REDIS DONE")

    // 2️⃣ generate reply (ตอนนี้ mock ไว้ก่อน)
    var reply string

    if msg == "ที่เที่ยว" {
		// call LLM API
        reply = "กาฬสินธุ์มี ภูสิงห์ + เขื่อนลำปาว"
    } else {
		// call LLM API
        reply = "ฉันคือ chatbot กาฬสินธุ์"
    }
	log.Println("🔥 HIT CHAT AI REPLY:", reply)
    // 3️⃣ save assistant reply only (user already saved before cache check)
	log.Println("STEP DB SAVE START")
    repository.Save(sessionID, "assistant", reply)
	log.Println("STEP DB SAVE DONE")
	
	log.Println("🔥 HIT CHAT AI SAVE DB")
    // 4️⃣ save cache
	log.Println("STEP REDIS SET START")
    err = config.RDB.Set(config.Ctx, msg, reply, 0).Err()
    if err != nil {
        log.Println("Redis set error:", err)
    }
    log.Println("STEP REDIS SET DONE")
	log.Println("🔥 HIT CHAT AI SAVE CACHE")
    return reply, nil
}