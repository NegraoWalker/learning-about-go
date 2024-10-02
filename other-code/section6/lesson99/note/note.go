package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct { //declaração da struct Note
	Title     string    `json:"title"`      //structs tags
	Content   string    `json:"content"`    //structs tags
	CreatedAt time.Time `json:"created_at"` //structs tags
}

func New(title, content string) (Note, error) { //função construtora usada para inicializar uma instância da struct Note

	if title == "" || content == "" {
		return Note{}, errors.New("error: invalid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() { //receiver para associar o método a struct Note
	fmt.Printf("Your note titled %v has the following content: %v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, error := json.Marshal(note)

	if error != nil {
		return error
	}

	return os.WriteFile(fileName, json, 0644)
}
