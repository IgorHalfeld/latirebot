package structs

type SendPayload struct {
	Name      string
	Provider  string
	Photo     string `json:"photo"`
	Caption   string `json:"caption"`
	ChatID    string `json:"chat_id"`
	ParseMode string `json:"parse_mode"`
}

type NotificationPayload struct {
	User          User
	Product       Product
	Caption       string
	NormalPrice   float64
	DiscountPrice float64
}
