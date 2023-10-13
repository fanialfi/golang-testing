package main

import "testing"

var (
	kubus              Kubus   = Kubus{4}
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

func TestHitungVolume(t *testing.T) {
	// method t.Logf() digunakan untuk memunculkan log,
	// method ini equivalen dengan fmt.Printf()
	t.Logf("Volume\t: %.2f", kubus.Volume())

	if kubus.Volume() != volumeSeharusnya {
		// method t.Errorf() digunakan untuk memunculkan log diikuti dengan keterangan bahwa terjadi fail pada saat testing berjalan
		t.Errorf("SALAH! harusnya %.2f", volumeSeharusnya)
	}
}

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas\t: %.2f", kubus.Luas())

	if kubus.Luas() != luasSeharusnya {
		t.Errorf("SALAH! harusnya %.2f", luasSeharusnya)
	}
}

func TestHitungKeliling(t *testing.T) {
	t.Logf("Keliling\t: %.2f", kubus.Keliling())

	if kubus.Keliling() != kelilingSeharusnya {
		t.Errorf("SALAH! harusnya %.2f", kelilingSeharusnya)
	}
}

// file untuk keperluan testing harus dipisah dengan file utama,
// namanya harus berakhiran _test.go, dan nama package nya harus sama.
// unit test di go dituliskan dalam bentuk function,
// yang memiliki parameter bertipe *testing.T dengan nama function harus diawali dengan nama Test
//
// untuk mengeksekusi testing adalah dengan menggunakan perintah `go test`
// disini karena struct yang akan diuji ada pada file kubus.go
// maka pada saat eksekusi nama file kubus.go dan kubus_test.go harus ditulis sebagai argument
// argument -v atau verbose digunakan untuk menampilkan semua output log pada saat pengujian

// Benchmark

func BenchmarkHitungLuas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.Luas()
	}
}

// package testing selain digunakan untuk unit test, juga berisikan tools untuk benchmarking
// cara pembuatannya dengan membuat function yang namanya diawali dengan Benchmark dan parameter-nya bertipe *testing.B
// untuk menjalankan Benchmark jalankan test dengan menambahkan	argument -bench=. / -bench namaBenchmark
// argument ini digunakan untuk menandai bahwa selain testing ada juga Benchmark yang perlu di uji
//
// hasil dari benchmarking akan terlihat seperti berikut ini
// BenchmarkHitungLuas-4           15654964                76.67 ns/op
// arti dari 15654964 adalah function di atas di tes sebanyak 15654964
// dan 76.67 ns/op berarti setiap satu kali run function membutuhkan waktu eata rata 76.67 nano detik
