package arithmetic_operations

func BasicOperations(angka1 int, angka2 int, jenis_operasi string) int {
	var hasil int
	switch jenis_operasi {
	case "pertambahan":
		hasil = angka1 + angka2
	case "pengurangan":
		hasil = angka1 - angka2
	case "perkalian":
		hasil = angka1 * angka2
	case "pembagian":
		hasil = angka1 / angka2
	}
	return hasil
}
