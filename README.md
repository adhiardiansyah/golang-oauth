Penjelasan alur aplikasi:

1. User mengakses halaman utama ("/") untuk menampilkan tombol sign in with google, kemudian user menekan tombol tersebut yang akan diarahkan ke choose an account, setelah itu user memilih salah satu akun google nya.

2. Dilakukan autentikasi, jika berhasil akan diarahkan ke endpont "/callback". Selanjutnya akan ada respon json yang diberikan, ada 4 kolom yaitu uuid, email, is_verified, dan picture.

3. Kemudian dilakukan pengecekan, apakah email user sudah ada di database ? jika sudah maka langsung dilakukan proses login, jika belum maka data user tersebut akan dimasukkan ke dalam database lalu dilanjutkan proses login dan akan mengembalikan respon json yang berisi data user tersebut ditambah token jwt dengan payload data berisi uuid.

4. Setelah proses sign in berhasil, user dapat mengakses data profil yang melalui end point "/user". Tentunya endpoint ini terdapat middleware dengan menyertakan Authorization Bearer "token" pada request header.

Dockerize dan implement CI/CD : docker pull adhiardiansyah/golang-oauth
