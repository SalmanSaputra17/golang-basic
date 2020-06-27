package main

import "fmt"

// use reflect to get type of variable
// use goto to jump to specific label

func printLine() {
	fmt.Println("----------------------------------------------")
}

func printList(data []map[string]string) {
	fmt.Println("Berikut daftar siswa RPL: ")
	for i, row := range data {
		fmt.Printf("%d.", i+1)
		fmt.Printf("Nama: %s\n  Usia: %s\n  Rombel: %s\n", row["nama"], row["usia"], row["rombel"])
	}
}

func main() {

	var dataSiswa = []map[string]string{
		{
			"nama":   "Salman Saputra",
			"usia":   "17",
			"rombel": "RPL XII-1",
		},
		{
			"nama":   "John Doe",
			"usia":   "17",
			"rombel": "RPL XII-2",
		},
		{
			"nama":   "Janny Doe",
			"usia":   "17",
			"rombel": "RPL XII-3",
		},
	}

	printList(dataSiswa)
	printLine()

	fmt.Println("1. Tambah data siswa")
	fmt.Println("2. Ubah data siswa")
	fmt.Println("3. Hapus data siswa")

	printLine()

	var yourChoice int
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scanf("%d", &yourChoice)

	printLine()

	switch yourChoice {
	case 1:
		var isi = map[string]string{
			"nama":   "Carissa Howland",
			"usia":   "17",
			"rombel": "RPL XII-4",
		}

		dataSiswa = append(dataSiswa, isi)
		printList(dataSiswa)

	case 2:
		dataSiswa[1] = map[string]string{
			"nama":   "Carissa Putri",
			"usia":   "17",
			"rombel": "RPL XII-1",
		}

		printList(dataSiswa)
	case 3:
		dataSiswa[len(dataSiswa)-1] = map[string]string{}
		dataSiswa = dataSiswa[:len(dataSiswa)-1]

		printList(dataSiswa)
	}

}
