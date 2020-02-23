package structs

// SendPayload is body request for send messages
type SendPayload struct {
	Name      string
	Photo     string `json:"photo"`
	Caption   string `json:"caption"`
	ChatID    string `json:"chat_id"`
	ParseMode string `json:"parse_mode"`
}
