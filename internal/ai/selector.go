package ai

import "strings"

func SelectRelevant(history []map[string]string, query string) []map[string]string {
	var selected []map[string]string

	for _, h := range history {
		if strings.Contains(h["content"], query) {
			selected = append(selected, h)
		}
	}

	if len(selected) == 0 {
		return history
	}
	return selected
}