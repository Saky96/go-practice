package note

import (
	"encoding/json"
	"errors"
	"example/interface/utils"
	"fmt"
	"os"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) DisplayNote() {
	fmt.Println("The details of this note is below:")
	fmt.Println("Title: ", note.Title)
	fmt.Println("Content: ", note.Content)
}

func (note Note) ValidateNote() error {
	if note.Content == "" || note.Title == "" {
		return errors.New("note content or title is empty")
	}
	return nil
}

func (note Note) Save() error {

	fileName := utils.SanitizeFileName(note.Title) + ".json"

	jsonData, err := json.Marshal(note)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)

}

func New(title string, content string) *Note {
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
}
