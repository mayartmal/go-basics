package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/main/note"
	"example.com/main/todo"
)

func main() {
	typeCheckerPrint(1)
	typeCheckerPrint("LJKVN")
	typeCheckerPrint(1.1)
	printTrueAnyData(1)
	printTrueAnyData("Hellio")
	printTrueAnyData(1.1)
	noteTitle, noteContent := getNoteData()

	note1, err := note.New(noteTitle, noteContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(note1)
	if err != nil {
		return
	}

	todo1, err := todo.New(getTodoData())
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo1)
	if err != nil {
		return
	}

}

func getNoteData() (string, string) {
	noteTitle := getUserInput("Note title: ")
	noteContent := getUserInput("Note content: ")
	return noteTitle, noteContent
}

func getTodoData() string {
	todoText := getUserInput("Add your todo: ")
	return todoText
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

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputter interface {
	saver
	Display()
}

// type outputter interface {
// 	Save() error
// 	Display()
// }

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("ERROR DURING SAVING")
		return err
	}
	fmt.Println("Saving succided")
	return nil

}

func outputData(data outputter) error {
	data.Display()
	return saveData(data)
}

func printAnyData(value interface{}) {
	fmt.Println(value)
}

func printTrueAnyData(value any) {
	switch value.(type) {
	case int:
		fmt.Println("Your integer is: ", value)
	case string:
		fmt.Println("Your string is: ", value)
	default:
		fmt.Println(value)
	}

}

func typeCheckerPrint(value any) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Int ", intVal)
	} else {
		fmt.Println("Unknown ", intVal)
	}
	stVal, ok := value.(string)
	if ok {
		fmt.Println("String ", stVal)
	}
}
