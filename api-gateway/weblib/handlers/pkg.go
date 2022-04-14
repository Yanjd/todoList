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

func PanicTaskError(err error) {
	if err != nil {
		logging.Info(errors.New("taskService--" + err.Error()))
		panic(err)
	}
}
