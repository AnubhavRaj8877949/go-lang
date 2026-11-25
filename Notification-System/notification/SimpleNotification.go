package notification

type SimpleNotification struct {
	Text string
}

func NewSimpleNotification(text string) *SimpleNotification {
	return &SimpleNotification{Text: text}
}

func (s *SimpleNotification) GetContent() string {
	return s.Text
}
