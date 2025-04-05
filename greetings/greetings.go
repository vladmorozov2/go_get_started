package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name 12321")
	}

	message := fmt.Sprintf(randomForest(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomForest() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you< %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
