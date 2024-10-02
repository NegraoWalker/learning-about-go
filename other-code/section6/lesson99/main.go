package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note"
	"example.com/todo"
)

func main() {

	title, content := GetNoteData()
	todoText := GetUserInput("Todo text: ")

	userTodo, error := todo.New(todoText)

	if error != nil {
		fmt.Println(error)
		return
	}

	userNote, error := note.New(title, content)

	if error != nil {
		fmt.Println(error)
		return
	}

	userTodo.Display()

	error = userTodo.Save()

	if error != nil {
		fmt.Println("error: saving the todo failed")
		return
	}

	fmt.Println("saving the todo succeeded!")

	userNote.Display()
	error = userNote.Save()

	if error != nil {
		fmt.Println("error: saving the note failed")
		return
	}

	fmt.Println("saving the note succeeded!")
}

func GetUserInput(prompt string) string {

	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin) //para armazenar dados de entrada do teclado que são maiores, usa-se o bufio

	text, error := reader.ReadString('\n') //a leitura da entrada vai parar quando ocorrer uma quebra de linha

	if error != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func GetNoteData() (string, string) {

	title := GetUserInput("Note title: ")

	content := GetUserInput("Note content: ")

	return title, content
}
