package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note"
	"example.com/todo"
)

type saver interface { //definição de interface
	Save() error //definição da função que deve ser implementada
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

// begin Main----------------------------------------------------------------------------
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

	error = OutputData(userTodo)

	if error != nil {
		fmt.Println("error: saving the todo failed")
		return
	}

	error = OutputData(userNote)

	if error != nil {
		return
	}
}

// end Main--------------------------------------------------------------------------------------------

func OutputData(data outputtable) error {
	data.Display()
	return SaveData(data)
}
func SaveData(data saver) error { //função que implementa a interface
	error := data.Save()

	if error != nil {
		fmt.Println("error: saving the note failed")
		return error
	}

	fmt.Println("saving the note succeeded!")
	return nil
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
