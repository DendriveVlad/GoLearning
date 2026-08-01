package utils

import (
	"TelephoneBook/db"
	"fmt"
)

func ProcessFind(findOne bool) ([]db.Contact, bool, error) {
	var contacts []db.Contact
	fmt.Println("Чтобы вернуться назад введите \"0\"")
	fmt.Println("Введите имя или номер контакта:")
	rq, _ := WaitForInput(false)
	if rq == "0" {
		return nil, false, nil
	}
	contact, err := db.SelectByPhone(rq)
	if err != nil {
		fmt.Println("Произошла ошибка с выгрузкой данных")
		return nil, false, err
	}
	if contact.Phone != "" {
		contacts = append(contacts, contact)
	} else {
		contacts, err = db.SelectByName(rq)
		if err != nil {
			fmt.Println("Произошла ошибка с выгрузкой данных ")
			return nil, false, err
		}
	}

	if len(contacts) <= 1 || !findOne {
		return contacts, true, nil
	}

	fmt.Println("=======Найдены  контакты=======")
	for i, contact := range contacts {
		fmt.Println(i+1, contact.Name, ":", contact.Phone)
	}
	fmt.Println("===============================")

	fmt.Println("Чтобы вернуться назад введите \"0\"")
	fmt.Println("Выберите номер контакта:")
	for {
		_, rq := WaitForInput(true)
		if rq == 0 {
			return contacts, true, nil
		}
		if rq < 1 || rq > len(contacts) {
			fmt.Println("Введено не верное значение, повторите запрос:")
			continue
		}
		var singleContact []db.Contact
		singleContact = append(singleContact, contacts[rq-1])
		return singleContact, false, nil
	}
}
