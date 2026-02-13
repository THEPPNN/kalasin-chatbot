package ai

func GuardPrompt(msg string) string {

	rules := `
If you don't know the answer
say "ไม่แน่ใจ"

Do not invent facts.
Do not fabricate places.
`

	return rules + "\nUser: " + msg
}