package db

import _ "modernc.org/sqlite"

func SelectAll() ([]Contact, error) {
	rs, err := DB.Query("SELECT * FROM " + TableName + " ORDER BY name")
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

func SelectByPhone(phone string) (Contact, error) {
	rs, err := DB.Query("SELECT * FROM "+TableName+" WHERE phone = ? ORDER BY name", phone)
	if err != nil {
		return Contact{}, err
	}
	defer rs.Close()
	var c Contact
	for rs.Next() {
		err := rs.Scan(&c.Phone, &c.Name)
		if err != nil {
			return c, err
		}
	}

	if err = rs.Err(); err != nil {
		return c, err
	}

	return c, nil
}

func SelectByName(name string) ([]Contact, error) {
	rs, err := DB.Query("SELECT * FROM "+TableName+" WHERE LOWER(name) LIKE ? ORDER BY name", "%"+name+"%")
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

func Update(oldPhone string, phone string, name string) error {
	_, err := DB.Exec("UPDATE "+TableName+" SET name = ?, phone = ? WHERE phone = ?", name, phone, oldPhone)
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
