package commands

import (
	"TelephoneBook/db"
	"TelephoneBook/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func DropData() bool {
	fmt.Println("После удаления данных програма завершит своё выполнение. Продолжить?")
	fmt.Println(`Напишите "Да" или "Нет":`)

	for {
		rq, _ := utils.WaitForInput(false)
		switch strings.ToLower(rq) {
		case "yes", "да":
			if err := dropDataDirectory(); err != nil {
				fmt.Println("Ошибка при удалении:", err)
				return false
			}
			fmt.Println("Данные удалены! До встречи!")
			return true
		case "0", "no", "нет":
			return true
		}
	}
}

func dropDataDirectory() error {
	if db.DB != nil {
		if err := db.DB.Close(); err != nil {
			return err
		}
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	return os.RemoveAll(filepath.Join(configDir, "TelephoneBook"))
}
