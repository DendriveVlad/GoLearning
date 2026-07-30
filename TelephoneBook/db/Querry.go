package db

import _ "modernc.org/sqlite"

func SelectAll() ([]Contact, error) {
	rs, err := DB.Query("SELECT * FROM " + TableName)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var contacts []Contact
	for rs.Next() {
		var c Contact
		err := rs.Scan(&c.Phone, &c.Name)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}

	if err = rs.Err(); err != nil {
		return nil, err
	}

	return contacts, nil
}

func Insert(phone string, name string) error {
	_, err := DB.Exec("INSERT INTO "+TableName+" (phone, name) VALUES (?, ?)", phone, name)
	if err != nil {
		return err
	}
	return nil
}

func UpdateName(phone string, name string) error {
	_, err := DB.Exec("UPDATE "+TableName+" SET name = ? WHERE phone = ?", name, phone)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePhone(phone string, newPhone string) error {
	_, err := DB.Exec("UPDATE "+TableName+" SET phone = ? WHERE phone = ?", newPhone, phone)
	if err != nil {
		return err
	}
	return nil
}

func Delete(phone string) error {
	_, err := DB.Exec("DELETE FROM "+TableName+" WHERE phone = ?", phone)
	if err != nil {
		return err
	}
	return nil
}
