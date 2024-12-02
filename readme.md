# Anime Management API

API untuk mengelola data anime menggunakan Go, Gin, dan GORM.

## Instalasi

1. Clone repositori ini:

   ```sh
   git clone https://github.com/username/anime-management-api.git
   cd anime-management-api

   ```

2. Connection to Database
   Sesuaikan pembuatan dengan .env, menggunakan MySQL

3. Migration Database

   **Up Migration**
   Perintah ini akan menerapkan semua migrasi yang belum diterapkan ke database Anda. Ini berguna ketika Anda ingin memperbarui skema database ke versi terbaru.

   ```
   migrate -database "mysql://root:@tcp(127.0.0.1:3306)/db_gorm" -path database/migrations up
   ```

   **Down Migration**
   Perintah ini akan membatalkan migrasi terakhir yang diterapkan. Ini berguna ketika Anda ingin mengembalikan skema database ke versi sebelumnya.

   ```
   migrate -database "mysql://root:@tcp(127.0.0.1:3306)/db_gorm" -path database/migrations down
   ```

## Menjalankan Aplikasi

1. Jalankan aplikasi:

```sh
go run main.go
```
