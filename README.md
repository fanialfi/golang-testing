# testing

go menyediakan package `testing` berisikan banyak tool untuk keperluan unti test.

table berikut berisikan method standar testing yang bisa digunakan di go

|method|kegunaan|
|------|--------|
|`Log()`|menampilkan log|
|`Logf()`|menampilkan log dengan format|
|`Fail()`|menandakan terjadi `Fail()` dan proses testing fungsi akan tetap diteruskan|
|`FailNow()`|menandakan terjadi `Fail()` dan proses testing fungsi dihentikan|
|`Failed()`|menampilkan laporan failed|
|`Error()`|`Log()` dengan diikuti `Fail()`|
|`Errorf()`|`Logf()` dengan diikuti `Fail()`|
|`Fatal()`|`Log()` dengan diikuti `FailNow()`|
|`Fatalf()`|`Logf()` dengan diikuti `FailNow()`|
|`Skip()`|`Log()` dengan diikuti `SkipNow()`|
|`Skipf()`|`Logf()` dengan diikuti `SkipNow()`|
|`SkipNow()`|menghentikan proses testing fungsi, dilanjutkan ke testing berikutnya|
|`Skiped()`|menampilkan laporan skip|
|`Parallel()`|meng-set bahwa eksekusi testing adalah parallel|

didalam package testify terdapat 5 package dengan kegunaan berbeda beda satu dengan yang lainnya.
Detailnya bisa dilihat pada table berikut ini.

|package|kegunaan|
|-------|--------|
|`assert`|berisikan tools standar untuk melakukan testing|
|`http`|berisikan tools untuk keperluan testing http|
|`mock`|berisikan tools untuk mocking object|
|`require`| sama seperti `assert` hanya saja jika terjadi fail pada saat test akan menghentikan eksekusi program|
|`suite`|berisikan tools testing yang berhubungan dengan struct dan method|