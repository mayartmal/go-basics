package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	// "strings"
	// "time"
)

// const FILENAME = "notes.txt"

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error) {
	text, err := promtValidation(text)

	if err != nil {
		return Todo{}, err
	}

	return Todo{
		Text: text,
	}, nil
}

func (t Todo) Display() {
	fmt.Printf("Your todo is: %v\n", t.Text)
}

func (t Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func promtValidation(text string) (string, error) {
	if text == "" {
		return text, errors.New("EMPTY CONTENT PROVIDED")
	}
	return text, nil
}
