package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/main/note"
)

func main() {
	noteTitle, noteContent := getNoteData()

	note1, err := note.New(noteTitle, noteContent)
	if err != nil {
		fmt.Println(err)
		return
	}
	note1.Display()
	err = note1.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}
	fmt.Println("Saving the note succided")

}

func getNoteData() (string, string) {
	noteTitle := getUserInput("Note title: ")
	noteContent := getUserInput("Note content: ")
	return noteTitle, noteContent
}

func getUserInput(askingPromt string) string {
	fmt.Print(askingPromt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text

}
