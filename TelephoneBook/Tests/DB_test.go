package Tests

import (
	"TelephoneBook/db"
	"errors"
	"testing"

	"modernc.org/sqlite"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]string
		want  int
	}{
		{
			name:  "Default Insert",
			input: map[string]string{"+79991234567": "Petya"},
			want:  0,
		}, {
			name:  "Duplicate Insert",
			input: map[string]string{"+79991234567": "Petya"},
			want:  1555,
		}, {
			name:  "Special number",
			input: map[string]string{"*222#": "Who am i"},
			want:  0,
		}, {
			name:  "New number",
			input: map[string]string{"+79291564569": "Petya"},
			want:  0,
		}, {
			name:  "Duplicate number with another name",
			input: map[string]string{"+79291564569": "Volodya"},
			want:  1555,
		},
	}

	db.TableName = "contacts_test"
	err := db.InitDB()
	if err != nil {
		panic(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.input {
				got := db.Insert(k, v)

				var sqliteErr *sqlite.Error

				if got != nil && errors.As(got, &sqliteErr) {
					if sqliteErr.Code() != tt.want {
						t.Errorf(
							"got %d, want %d",
							sqliteErr.Code(),
							tt.want,
						)
					}
				}
			}
		})
	}
	_, err = db.DB.Exec("DROP TABLE contacts_test")
	if err != nil {
		panic(err)
	}
}
