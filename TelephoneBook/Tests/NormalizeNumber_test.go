package Tests

import (
	"TelephoneBook/utils"
	"testing"
)

func TestNormalizeNumber(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "+79991234567",
			want:  "+79991234567",
		},
		{
			input: "89991234567",
			want:  "89991234567",
		},
		{
			input: "*201#",
			want:  "*201#",
		},
		{
			input: "*#*201*#*",
			want:  "*#*201*#*",
		},
		{
			input: "*201",
			want:  "",
		},
		{
			input: "ABOBA",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {

			got, _ := utils.NormalizeNumber(tt.input)

			if got != tt.want {
				t.Errorf(
					"got %s, want %s",
					got,
					tt.want,
				)
			}
		})
	}

}
