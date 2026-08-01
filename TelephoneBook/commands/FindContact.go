package commands

import (
	"TelephoneBook/utils"
	"fmt"
)

func FindContact() bool {
	for {
		contacts, repeat, err := utils.ProcessFind(false)
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
			fmt.Println("=======Найденный контакт=======")
			fmt.Println(contacts[0].Phone, ":", contacts[0].Name)
			fmt.Println("===============================")
			continue
		}
		fmt.Println("=======Найдены  контакты=======")
		for _, contact := range contacts {
			fmt.Println(contact.Name, ":", contact.Phone)
		}
		fmt.Println("===============================")
	}
	return true
}
