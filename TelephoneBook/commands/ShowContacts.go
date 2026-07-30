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
	fmt.Println("====Контакты из справочника====")
	for _, contact := range contacts {
		fmt.Println(contact.Phone, ":", contact.Name)
	}
	fmt.Println("===============================")
	return true
}
