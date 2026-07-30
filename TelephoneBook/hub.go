package main

import (
	"TelephoneBook/commands"
	"TelephoneBook/utils"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func hub() error {
Loop:
	for {
		fmt.Println("Выберите команду:")
		fmt.Println("1. Добавить контакт")
		fmt.Println("2. Найти контакт")
		fmt.Println("3. Удалить контакт")
		fmt.Println("4. Редактировать контакт")
		fmt.Println("5. Показать все контакты")
		fmt.Println("9. Отчистить контакты (удалить все данные)")
		fmt.Println("0. Выход")

		_, cmd := utils.WaitForInput(true)
		fmt.Println()
		switch cmd {
		case 1:
			commands.AddContact()
		case 2:
		case 3:
		case 4:
		case 5:
			commands.ShowContacts()
		case 9:
		case 0:
			break Loop
		default:
			fmt.Println("Не известная команда")
		}
		fmt.Println()
	}
	return nil
}
