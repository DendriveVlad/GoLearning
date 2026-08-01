package commands

import (
	"TelephoneBook/db"
	"TelephoneBook/utils"
	"fmt"
	"strings"
)

func DropContact() bool {
	for {
		contacts, repeat, err := utils.ProcessFind(true)
		if err != nil {
			return false
		}
		if contacts == nil && !repeat {
			break
		}
		if (contacts == nil || len(contacts) == 0) && repeat {
			fmt.Println("=======Ничего не найдено=======")
			continue
		}
		if len(contacts) == 1 && contacts[0].Phone != "" {
			ok := dropContact(contacts[0])
			if !ok {
				return false
			}
		}
		continue
	}
	return true
}

func dropContact(contact db.Contact) bool {
Loop:
	for {
		fmt.Println("Вы уверене, что хотите удалить контакт:", contact.Phone, ":", contact.Name+"?")
		fmt.Println("Напишите \"Да\" или \"Нет\":")
		rq, _ := utils.WaitForInput(false)
		switch strings.ToLower(rq) {
		case "yes", "да":
			err := db.Delete(contact.Phone)
			if err != nil {
				fmt.Println("Произошла ошибка с удалением данных")
				return false
			}
			fmt.Println("Контакт удалён!")
			fmt.Println()
			fallthrough

		case "0", "no", "нет":
			break Loop
		}
	}
	return true
}
