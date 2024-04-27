API DATA DOSEN

Ini adalah API sederhana untuk mengelola data tentang dosen (dosen) di suatu universitas. API memungkinkan kita untuk melakukan operasi CRUD pada data, termasuk membuat, membaca, memperbarui, dan menghapus catatan.

Titik akhir
API memiliki titik akhir berikut:

GET /dosen: Mengambil daftar semua catatan dosen.
GET /dosen/:id: Mengambil satu catatan dosen berdasarkan ID-nya.
POST /dosen: Membuat catatan dosen baru.
PUT /dosen/:id: Memperbarui catatan dosen yang ada berdasarkan ID-nya.
DELETE /dosen/:id: Menghapus record dosen berdasarkan ID-nya.
Format Permintaan dan Tanggapan
API menggunakan JSON sebagai format permintaan dan respons. 

Berikut ini contoh tampilan permintaan dan respons:

REQUEST json :

{
  "nama_lengkap": "John Doe",
  "nip": "1234567890",
  "jenis_kelamin": "Laki-laki",
  "tempat_lahir": "Jakarta",
  "tanggal_lahir": "2000-01-01",
  "alamat": "Jl. Merdeka No. 1",
  "no_hp": "081234567890",
  "fakultas": "Teknik Informatika"
}

RESPONSE json :

{
  "id": 1,
  "nama_lengkap": "John Doe",
  "nip": "1234567890",
  "jenis_kelamin": "Laki-laki",
  "tempat_lahir": "Jakarta",
  "tanggal_lahir": "2000-01-01",
  "alamat": "Jl. Merdeka No. 1",
  "no_hp": "081234567890",
  "fakultas": "Teknik Informatika"
}

Dalam contoh ini, permintaannya adalah permintaan POST ke endpoint /dosen, dan responsnya adalah rekaman dosen yang baru dibuat.

Error Handling
API mengembalikan kode status HTTP yang sesuai untuk kesalahan, seperti 400 Bad Request untuk permintaan yang tidak valid, 401 Not Valid untuk permintaan tidak sah, dan 404 Not Found untuk permintaan ke sumber daya yang tidak ada.

Autentikasi
API menggunakan basic authentification untuk semua endpoint

