package ai

func BuildPrompt(
	system string,
	history []map[string]string,
	userMsg string,
	maxTokens int,
) []Message {

	var messages []Message

	// system
	messages = append(messages, Message{
		Role: "system",
		Content: system,
	})

	// select relevant memory
	history = SelectRelevant(history, userMsg)

	// add history
	for i := len(history) - 1; i >= 0; i-- {
		messages = append(messages, Message{
			Role: history[i]["role"],
			Content: history[i]["content"],
		})
	}

	// guard + user msg
	messages = append(messages, Message{
		Role: "user",
		Content: GuardPrompt(userMsg),
	})

	// trim tokens
	messages = TrimMessages(messages, maxTokens)

	return messages
}