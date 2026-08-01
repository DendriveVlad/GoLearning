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
		ok := true
		switch cmd {
		case 1:
			ok = commands.AddContact()
		case 2:
			ok = commands.FindContact()
		case 3:
			ok = commands.DropContact()
		case 4:
			ok = commands.UpdateContact()
		case 5:
			ok = commands.ShowContacts()
		case 9:
			ok = commands.DropData()
			if ok {
				break Loop
			}
		case 0:
			break Loop
		default:
			fmt.Println("Не известная команда")
		}
		if ok == false {
			fmt.Println("При работе команды что-то пошло не так...")
		}
		fmt.Println()
	}
	return nil
}
