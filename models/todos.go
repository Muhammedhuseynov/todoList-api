package models

import "github.com/Muhammedhuseynov/ginAPItodo_app/configData"

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (t *Todo) TableName() string {
	return "todo"
}

func GetAllTodos(todoList *[]Todo) (err error) {
	if err = configData.DB.Find(todoList).Error; err != nil {
		return err
	}
	return nil
}

func GetATodo(todoList *Todo, id string) (err error) {
	if err := configData.DB.Where("id = ?", id).First(todoList).Error; err != nil {
		return err
	}
	return nil
}

func CreateTodoList(todoList *Todo) (err error) {
	if err = configData.DB.Create(todoList).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodoList(todoList *Todo, id string) (err error) {
	configData.DB.Save(todoList)
	return nil
}

func DeleteToListItem(todoList *Todo, id string) (err error) {
	configData.DB.Where("id = ?", id).Delete(todoList)
	return nil
}
