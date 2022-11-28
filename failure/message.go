package failure

type ErrorMessage struct {
	Message string `json:"message"`
}

func Msg(message string) *ErrorMessage {
	return &ErrorMessage{Message: message}
}
