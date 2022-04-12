package handlers

import (
	"api-gateway/pkg/logging"
	"errors"
)

func PanicUserError(err error) {
	if err != nil {
		logging.Info(errors.New("userService--" + err.Error()))
		panic(err)
	}
}
