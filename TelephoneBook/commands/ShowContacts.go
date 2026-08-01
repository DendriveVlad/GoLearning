package commands

import (
	"TelephoneBook/db"
	"fmt"
)

func ShowContacts() bool {
	contacts, err := db.SelectAll()
	if err != nil {
		fmt.Println("Произошла ошибка с выгрузкой данных")
		return false
	}
	if contacts == nil || len(contacts) == 0 {
		fmt.Println("Справочник пуст")
		return true
	}
	fmt.Println("====Контакты из справочника====")
	for _, contact := range contacts {
		fmt.Println(contact.Name, ":", contact.Phone)
	}
	fmt.Println("===============================")
	return true
}
