package status

import "fmt"

type MyError struct {
	Msg string
}

func (err *MyError) Error() string {
	return fmt.Sprintf("ERROR: %s", err.Msg)
}
