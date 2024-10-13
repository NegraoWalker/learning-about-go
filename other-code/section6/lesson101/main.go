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

	printSomething(1)
	printSomething(2.89645)
	printSomething("Text")
	printSomething(false)

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
// Form 1:
// func printSomething(value interface{}) { //a sintaxe interface{} permite que a função receba qualquer tipo de valor como parâmetro
//
//		switch value.(type) { //sintaxe switch value.(type) permite criar ações para cada tipo de dado
//		case int:
//			fmt.Println("Integer:", value)
//		case float64:
//			fmt.Println("Float:", value)
//		case string:
//			fmt.Println("String:", value)
//		default:
//			fmt.Println("Other:", value)
//		}
//	}
//
// Form 2:
func printSomething(value interface{}) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", intVal)
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float:", floatVal)
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("Integer:", stringVal)
	}
}

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
