package entity

type Message struct {
	ChatID    int64
	MessageID *int
	Text      string
	Markup    *string
}
