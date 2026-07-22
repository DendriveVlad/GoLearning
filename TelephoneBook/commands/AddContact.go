package commands

import (
	"TelephoneBook/utils"
	"fmt"
)

func AddContact() bool {
	fmt.Println("Чтобы вернуться назад введите: 0")
	fmt.Println("Введите имя контакта:")
	name, _ := utils.WaitForInput(false)
	if name == "0" {
		return false
	}

	for {
		fmt.Println("Введите номер телефона:")
		number, _ := utils.WaitForInput(false)
		if number == "0" {
			return false
		}
		number, ok := utils.NormalizeNumber([]rune(number))
		if !ok {
			fmt.Println("Номер не распознан, попробуйте снова")
		}
		break
	}

	return true
}
