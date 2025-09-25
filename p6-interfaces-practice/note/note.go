package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// const FILENAME = "notes.txt"

type Note struct {
	NoteTitle   string    `json:"note_title"`
	NoteContent string    `json:"note_content"`
	CreatedAt   time.Time `json:"created_at"`
}

func New(noteTitle, noteContent string) (Note, error) {
	noteTitle, noteContent, err := promtValidation(noteTitle, noteContent)

	if err != nil {
		return Note{}, err
	}

	return Note{
		NoteTitle:   noteTitle,
		NoteContent: noteContent,
		CreatedAt:   time.Now(),
	}, nil
}

func (n Note) Display() {
	fmt.Printf("Your note title is: %v\n", n.NoteTitle)
	fmt.Printf("Your note content is: %v\n", n.NoteContent)
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.NoteTitle, " ", "-")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(n)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func promtValidation(title, content string) (string, string, error) {
	if title == "" || content == "" {
		return title, content, errors.New("EMPTY CONTENT PROVIDED")
	}
	return title, content, nil
}
