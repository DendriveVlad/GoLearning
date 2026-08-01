package commands

import (
	"TelephoneBook/db"
	"TelephoneBook/utils"
	"fmt"
	"strings"
)

func UpdateContact() bool {
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
			fmt.Println("Изменение контакта:", contacts[0].Name, ":", contacts[0].Phone)
			fmt.Println("Чтобы вернуться назад введите \"0\"")
			newContact := db.Contact{}
		Start:
			for {
				fmt.Println("Новое имя контакта (чтобы не менять напишите \"-\"):")
				name, _ := utils.WaitForInput(false)
				if name == "0" {
					break
				} else if name == "-" {
					name = contacts[0].Name
				}
				fmt.Println("Чтобы вернуться назад введите \"0\"")
				fmt.Println("Новый номер контакта (чтобы не менять напишите \"-\"):")
				for {
					number, _ := utils.WaitForInput(false)
					if number == "0" {
						break
					} else if number == "-" {
						number = contacts[0].Phone
					}
					number, ok := utils.NormalizeNumber(number)
					if !ok {
						fmt.Println("Номер не распознан, попробуйте снова")
						continue
					}
					if number == contacts[0].Phone && name == contacts[0].Name {
						fmt.Println("Нет изменений")
						break Start
					}
					newContact.Phone = number
					newContact.Name = name
					_ = updateContact(contacts[0], newContact)
					break Start
				}
			}
		}
		continue
	}
	return true
}

func updateContact(oldContact db.Contact, newContact db.Contact) bool {
Loop:
	for {
		fmt.Println("Контакт будет изменён:")
		fmt.Println(oldContact.Name, ":", oldContact.Phone)
		fmt.Println("↓")
		fmt.Println(newContact.Name, ":", newContact.Phone)
		fmt.Println("Вы согласны? Напишите \"Да\" или \"Нет\":")
		rq, _ := utils.WaitForInput(false)
		switch strings.ToLower(rq) {
		case "yes", "да":
			err := db.Update(oldContact.Phone, newContact.Phone, newContact.Name)
			if err != nil {
				fmt.Println("Произошла ошибка с изменением данных")
				return false
			}
			fmt.Println("Контакт изменён!")
			fmt.Println()
			fallthrough

		case "0", "no", "нет":
			break Loop
		}
	}
	return true
}
