package main

import (
	"example/structs/admin"
	"example/structs/note"
	"example/structs/user"
	"example/structs/utils"
	"fmt"
)

func main() {
	noteMain()
}

func userMain() {
	userFirstName := user.GetUserData("Please enter your first name: ")
	userLastName := user.GetUserData("Please enter your last name: ")
	userBirthdate := user.GetUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.NewCreateUserValidation(userFirstName, userLastName, userBirthdate)

	if err != nil {
		return
	}

	// ... do something awesome with that gathered data!
	// var userData user
	// userData := user.User{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthdate,
	// 	createdAt: time.Now(),
	// }
	// userData.OutputUserDataStructMethod()
	// outputUserData(userData) //normally using struct data in a function
	// outputUserDataUsingPointer(&userData) // pointer to structs are passed
	// userData.clearUserName()
	// userData.outputUserDataStructMethod()
	//appUser.OutputUserDataStructMethod()
	//userData2 := user.New(
	//	"Shaurya",
	//	"Saigal",
	//	"17/02/2004",
	//)
	//
	//userData2.OutputUserDataStructMethod()
	//
	//userData3 := user.NewCreateUserPointer(
	//	"avi",
	//	"Saigal",
	//	"22/11/2005",
	//)
	//
	//userData3.OutputUserDataStructMethod()

	adminPointer := admin.NewAdmin(
		"admin@email.com",
		"*****",
		*appUser,
	)

	adminValue := *adminPointer

	fmt.Println(adminValue.Email())
	adminPointer.OutputUserDataStructMethod()
}

// Note struct to store note information
//type Note struct {
//	title   string
//	content string
//}
//
//// DisplayNote Method to display note details
//func (n *Note) DisplayNote() {
//	fmt.Println("The details of this note are below:")
//	fmt.Println("Title: ", n.title)
//	fmt.Println("Content: ", n.content)
//}
//
//// Function to get user input
//func getUserInput(prompt string) string {
//	reader := bufio.NewReader(os.Stdin)
//	fmt.Print(prompt)
//	input, _ := reader.ReadString('\n')
//	// Remove trailing newline characters
//	return strings.TrimSpace(input)
//}

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
	fileErr := note1.SaveNote()
	if err != nil {
		fmt.Println(fileErr)
		//return err
	}

	fmt.Println("Note successfully saved!")

}
