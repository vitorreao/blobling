package failure

type ErrorMessage struct {
	Message string `json:"message"`
}

func NewErrorMsg(message string) *ErrorMessage {
	return &ErrorMessage{Message: message}
}
