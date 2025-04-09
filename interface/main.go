package main

import (
	"example/interface/note"
	"example/interface/todo"
	"example/interface/utils"
	"fmt"
)

type Saver interface {
	Save() error
}

func main() {
	noteMain()
}

func noteMain() {
	fmt.Println("Add a note!")
	// Create a new Note
	//newNote := &Note{}

	// Get title from user
	title := utils.GetUserInput("Enter the note title: ")

	// Get content from user
	content := utils.GetUserInput("Enter the note content: ")

	//newNote := &note.Note{}

	// Display the note
	//newNote.DisplayNote(title, content)

	note1 := note.New(title, content)
	err := note1.ValidateNote()
	if err != nil {
		fmt.Println(err)
		return
	}

	note1.DisplayNote()
	fmt.Println("Saving the note!")

	err = saveData(note1)

	fmt.Println("Note successfully saved!")

	item := utils.GetUserInput("Enter the item title: ")
	todo1 := todo.New(item)

	err = saveData(todo1)
}

func saveData(data Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
