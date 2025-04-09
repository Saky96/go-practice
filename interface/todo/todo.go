package todo

import (
	"encoding/json"
	"os"
)

type Todo struct {
	Item string `json:"item"`
}

func New(item string) *Todo {
	return &Todo{
		Item: item,
	}
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	value, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, value, 0644)

}
