package ai

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func BuildPrompt(system string, history []map[string]string, userMsg string) []Message {

	var messages []Message

	// system prompt
	messages = append(messages, Message{
		Role: "system",
		Content: system,
	})

	// history
	for i := len(history) - 1; i >= 0; i-- {
		messages = append(messages, Message{
			Role: history[i]["role"],
			Content: history[i]["content"],
		})
	}

	// latest user msg
	messages = append(messages, Message{
		Role: "user",
		Content: userMsg,
	})

	return messages
}