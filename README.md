# testing

go menyediakan package `testing` berisikan banyak tool untuk keperluan unti test.

untuk membuat unit test file akhirannya harus berupa `_test.go`, 
selain itu aturan pembuatan function unit test nama function harus di awali dengan nama **Test** selain itu juga harus memiliki parameter `t *testing.T` dan tidak mengembalikan value.

untuk menjalankan file unit test bisa menggunakan perintah `go test -v` perintah tersebut akan menjalankan semua function testing yang ada di current directory.

untuk menjalankan spesifik function apa yang ingin dijalankan bisa dengan menggunakan perintah `go test -v -run ^TestNamaFunction$`

untuk menggagalkan unit test terdapat function `Fail()` dan `FailNow()`,
`Fail()` akan menggagalkan test, namun tetap melanjutkan eksekusi unit test, namun di akhir akan dianggap gagal,
`FailNow()` akan menggagalkan test saat ini juga dan tidak akan melanjutkan eksekusi test.

|method|keterangan|
|------|----------|
|`Fail()`|menandakan terjadi `Fail()` dan proses testing fungsi akan tetap diteruskan|
|`FailNow()`|menandakan terjadi `Fail()` dan proses testing fungsi dihentikan|


ketika menggunakan `Fail()` dan `FailNow()` kita tidak tahu penyebab gagalnya test, supaya kita tahu penyebab gagalnya test kita bisa menggunakan `Error()` atau `Errorf()` dan `Fatal()` atau `Fatalf()`,

`Error()` dan `Errorf()` seperti melakukan log error lalu memanggil `Fail()` namun dengan `Errorf()` kita bisa melakukan format untuk pesan log nya.
`Fatal()` atau `Fatalf()` seperti melakukan log error lalu memanggil `FailNow()` namun dengan `Fatalf()` kita bisa melakukan format untuk pesan log nya.

|method|keterangan|
|------|----------|
|`Log()`|menampilkan log|
|`Logf()`|menampilkan log dengan format|
|`Error()`|`Log()` dengan diikuti `Fail()`|
|`Errorf()`|`Logf()` dengan diikuti `Fail()`|
|`Fatal()`|`Log()` dengan diikuti `FailNow()`|
|`Fatalf()`|`Logf()` dengan diikuti `FailNow()`|

melakukan pengecekan di unit test sangat disarankan menggunakan assertion, 
golang tidak menyediakan package untuk melakukan assertion, 
sehingga perlu menggunakan package pihak ketiga untuk melakukanya, dan salah satunya adalah `testify`

`testify` mempunyai 2 package untuk melakuka assertion, yaitu `assert` dan `require`, 
jika menggunakan `assert` jika test gagal maka akan memanggil `Fail()`,
jika menggunakan `require` jika test gagal maka akan memanggil `FailNow()`

|package|kegunaan|
|-------|--------|
|`assert`|berisikan tools standar untuk melakukan testing|
|`require`| sama seperti `assert` hanya saja jika terjadi fail pada saat test akan menghentikan eksekusi program|

table berikut berisikan method standar testing yang bisa digunakan di go

|method|kegunaan|
|------|--------|
|`Failed()`|menampilkan laporan failed|
|`Skip()`|`Log()` dengan diikuti `SkipNow()`|
|`Skipf()`|`Logf()` dengan diikuti `SkipNow()`|
|`SkipNow()`|menghentikan proses testing fungsi, dilanjutkan ke testing berikutnya|
|`Skiped()`|menampilkan laporan skip|
|`Parallel()`|meng-set bahwa eksekusi testing adalah parallel|

didalam package testify terdapat 5 package dengan kegunaan berbeda beda satu dengan yang lainnya.
Detailnya bisa dilihat pada table berikut ini.

|package|kegunaan|
|-------|--------|
|`http`|berisikan tools untuk keperluan testing http|
|`mock`|berisikan tools untuk mocking object|
|`suite`|berisikan tools testing yang berhubungan dengan struct dan method|