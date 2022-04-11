package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returna a greetings for the named person
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprintf(randomFormat())
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages
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

// init sets initial value for variables use in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	//A slice of message formats
	formats := []string{
		"Hi, %v, Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
