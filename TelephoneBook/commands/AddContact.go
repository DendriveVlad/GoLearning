package commands

import (
	"TelephoneBook/db"
	"TelephoneBook/utils"
	"errors"
	"fmt"

	"modernc.org/sqlite"
)

func AddContact() bool {
	fmt.Println("Чтобы вернуться назад введите: 0")
	fmt.Println("Введите имя контакта:")
	var (
		name   string
		number string
	)
	name, _ = utils.WaitForInput(false)
	if name == "0" {
		return false
	}

	for {
		fmt.Println("Введите номер телефона:")
		number, _ = utils.WaitForInput(false)
		if number == "0" {
			return false
		}
		var ok bool
		number, ok = utils.NormalizeNumber(number)
		if !ok {
			fmt.Println("Номер не распознан, попробуйте снова")
			continue
		}
		break
	}

	err := db.Insert(number, name)
	if err != nil {
		var sqliteErr *sqlite.Error
		if errors.As(err, &sqliteErr) && sqliteErr.Code() == 1555 {
			fmt.Println("Такой номер телефона уже записан")
		} else {
			fmt.Println("Произошла ошибка при записи в бд")
		}
		return false
	}

	fmt.Println("Номер сохранён в справочнике!")
	return true
}
