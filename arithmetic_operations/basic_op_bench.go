package arithmetic_operations

import (
	"testing"
)

func BenchmarkBasicOperations(b *testing.B) {
	tests := []struct {
		angka1        int
		angka2        int
		jenis_operasi string
	}{
		{
			angka1:        6,
			angka2:        2,
			jenis_operasi: "perkalian",
		}, {
			angka1:        68,
			angka2:        29,
			jenis_operasi: "perkalian",
		}, {
			angka1:        6123,
			angka2:        2456,
			jenis_operasi: "perkalian",
		},
	}
	for _, test := range tests {
		b.Run(test.jenis_operasi, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BasicOperations(test.angka1, test.angka2, test.jenis_operasi)
			}
		})
	}
}
