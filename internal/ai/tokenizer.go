package ai

func EstimateTokens(text string) int {
	return len(text) / 4
}

func TrimMessages(messages []Message, maxTokens int) []Message {
	total := 0
	var result []Message

	for i := len(messages) - 1; i >= 0; i-- {
		t := EstimateTokens(messages[i].Content)
		if total+t > maxTokens {
			break
		}
		total += t
		result = append([]Message{messages[i]}, result...)
	}
	return result
}