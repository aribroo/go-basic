package main

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

type validationError struct {
	message    string
	codeStatus int
}

func TestCustomErrors(t *testing.T) {

	username := ""
	password := ""

	var vErr *validationError

	if result, err := login(username, password); err != nil {
		if errors.As(err, &vErr) {
			fmt.Println(err.Error())
		} else {
			fmt.Println(result)
		}
	}

}

func (v *validationError) Error() string {
	return v.message + " with status code " + strconv.Itoa(v.codeStatus)
}

func login(username, password string) (string, error) {
	if username == "" || password == "" {
		return "", &validationError{message: "username or password is wrong", codeStatus: 400}
	}

	return "login successfully", nil
}
