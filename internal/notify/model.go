package notify

type Message struct {
	To   string `json:"to"`
	Body string `json:"body"`
}