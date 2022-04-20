package arithmetic_operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicOperations(t *testing.T) {
	tests := []struct {
		angka1        int
		angka2        int
		jenis_operasi string
		expected      int
	}{
		{
			angka1:        6,
			angka2:        2,
			jenis_operasi: "pertambahan",
			expected:      8,
		},
		{
			angka1:        6,
			angka2:        2,
			jenis_operasi: "pengurangan",
			expected:      4,
		},
		{
			angka1:        6,
			angka2:        2,
			jenis_operasi: "perkalian",
			expected:      12,
		},
		{
			angka1:        6,
			angka2:        2,
			jenis_operasi: "pembagian",
			expected:      3,
		},
	}
	for _, test := range tests {
		t.Run(test.jenis_operasi, func(t *testing.T) {
			result := BasicOperations(test.angka1, test.angka2, test.jenis_operasi)
			assert.Equal(t, test.expected, result)
			assert.NotNil(t, result)
		})
	}
}
