# gRPC Polyglot Lab

Project ini merupakan implementasi komunikasi gRPC menggunakan bahasa pemrograman Go, Python, dan Node.js.

## 1. Persiapan Project

Buat folder project dengan nama:

`grpc-polyglot-lab`

Project ini terdiri dari folder `server` untuk program Go, folder `clients` untuk client Python dan Node.js, folder `protoc`, serta file `user.proto` untuk mendefinisikan layanan dan struktur data gRPC.

Buka folder project tersebut menggunakan Visual Studio Code.

## 2. Persiapan Perangkat Lunak

Perangkat lunak yang digunakan dalam project ini adalah:

- Go
- Python
- Node.js
- npm
- Protocol Buffers Compiler (protoc)
- Visual Studio Code

Untuk memastikan perangkat lunak sudah terpasang, jalankan perintah:

```bash
go version
python --version
node --version
npm --version
protoc --version
```

Jika versi masing-masing perangkat lunak ditampilkan, berarti instalasi berhasil.

## 3. Membuat File user.proto

File `user.proto` digunakan untuk mendefinisikan layanan gRPC dan struktur data pengguna.

Layanan yang dibuat adalah `UserService` dengan fungsi `GetUser`. Fungsi tersebut menerima ID pengguna dan mengembalikan data berupa ID, nama, dan email.

## 4. Membuat dan Menjalankan Server Go

Server dibuat menggunakan bahasa pemrograman Go.

Masuk ke folder server:

```bash
cd server
```

Jalankan server menggunakan perintah:

```bash
go run ./cmd/server
```

Server berjalan pada port:

`50051`

Jika berhasil, terminal akan menampilkan pesan:

```text
Server gRPC berjalan di port 50051
```

Biarkan server tetap berjalan selama pengujian client.

## 5. Menyiapkan Client Python

Client Python dibuat di dalam folder `clients/python`.

Library yang digunakan adalah `grpcio` dan `grpcio-tools`.

Aktifkan virtual environment yang sudah dibuat:

```powershell
.\venv\Scripts\Activate.ps1
```

Client Python menggunakan file `user.proto` untuk berkomunikasi dengan server Go melalui gRPC.

## 6. Menjalankan Client Python

Pastikan server Go masih berjalan.

Masuk ke folder `clients/python`, kemudian jalankan:

```powershell
python client.py
```

Jika berhasil, akan ditampilkan data pengguna berupa ID, nama, dan email.

Contoh hasil pengujian:

```text
Data pengguna berhasil diterima:
ID    : 1
Nama  : Nisatur Rifka
Email : nisaturrifka@gmail.com
```

## 7. Menyiapkan Client Node.js

Client Node.js dibuat di dalam folder `clients/node`.

Library yang digunakan adalah:

- `@grpc/grpc-js`
- `@grpc/proto-loader`

Library tersebut digunakan untuk melakukan komunikasi gRPC dan membaca definisi layanan dari file `user.proto`.

## 8. Menjalankan Client Node.js

Pastikan server Go masih berjalan.

Masuk ke folder `clients/node`, kemudian jalankan:

```bash
node client.js
```

Jika berhasil, client akan menampilkan data pengguna yang diterima dari server Go.

## 9. Memeriksa Hasil Pengujian

Pengujian dilakukan dengan menjalankan server Go, kemudian menjalankan client Python dan Node.js secara bergantian.

Server menerima permintaan dari kedua client melalui port `50051`.

Hasil pengujian menunjukkan bahwa client Python dan Node.js berhasil mengambil data pengguna dari server Go menggunakan layanan gRPC.

## 10. Kesimpulan

Project ini menunjukkan bahwa gRPC dapat digunakan untuk menghubungkan aplikasi yang dibuat menggunakan bahasa pemrograman berbeda.

Pada project ini, Go digunakan sebagai server, sedangkan Python dan Node.js digunakan sebagai client. Kedua client berhasil mengirim permintaan dan menerima data pengguna dari server melalui layanan `GetUser`.

Dengan demikian, komunikasi antara Go, Python, dan Node.js menggunakan protokol gRPC berhasil diterapkan.