package commands

import (
	"TelephoneBook/db"
	"TelephoneBook/utils"
	"errors"
	"fmt"

	"modernc.org/sqlite"
)

func AddContact() bool {
	contact := db.Contact{}

	fmt.Println("Чтобы вернуться назад введите \"0\"")
	fmt.Println("Введите имя контакта:")

	contact.Name, _ = utils.WaitForInput(false)
	if contact.Name == "0" {
		return true
	}

	for {
		fmt.Println("Введите номер телефона:")
		contact.Phone, _ = utils.WaitForInput(false)
		if contact.Phone == "0" {
			return true
		}
		var ok bool
		contact.Phone, ok = utils.NormalizeNumber(contact.Phone)
		if !ok {
			fmt.Println("Номер не распознан, попробуйте снова")
			continue
		}
		break
	}

	err := db.Insert(contact.Phone, contact.Name)
	if err != nil {
		var sqliteErr *sqlite.Error
		if errors.As(err, &sqliteErr) && sqliteErr.Code() == 1555 {
			fmt.Println("Такой номер телефона уже записан")
			return true
		}
		fmt.Println("Произошла ошибка при записи в бд")
		return false
	}

	fmt.Println("Номер сохранён в справочнике!")
	return true
}
