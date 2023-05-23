package greetings

import (
	"errors"
	"fmt"
)

func Hello1(name string) string {
	var message = "Hi, Welcome!"
	return message
}

func Hello2(name string) string {
	message := "Hi, Welcome!"
	return message
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
