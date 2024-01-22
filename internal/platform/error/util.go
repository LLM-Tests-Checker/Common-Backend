package error

import (
	"errors"
	"fmt"
)

func Wrap(err error, message string) error {
	return errors.New(fmt.Sprintf("%s: %s", message, err))
}
