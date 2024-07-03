package myerror

import "fmt"

// Custom interface error
type ErrorInvalidArgument interface {
	Error() string
}

type errorInvalidArgument struct {
	Msg string
}

func (e errorInvalidArgument) Error() string {
	return fmt.Sprintf("invalid argument: %s", e.Msg)
}

func NewErrorInvalidArgument(msg string) ErrorInvalidArgument {
	// Implement interface MyError
	return errorInvalidArgument{Msg: msg}
}
