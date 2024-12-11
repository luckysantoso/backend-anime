# Anime Management API

API untuk mengelola data anime menggunakan Go, Gin, dan GORM.

## Instalasi

1. Clone repositori ini:

   ```sh
   git clone https://github.com/luckysantoso/backend-anime.git

   ```

2. Connection to Database
   Sesuaikan pembuatan dengan .env, menggunakan MySQL atau PostgreSQL

3. Migration Database

   a. Migration Database With MySQL
   **Up Migration**
   Perintah ini akan menerapkan semua migrasi yang belum diterapkan ke database Anda. Ini berguna ketika Anda ingin memperbarui skema database ke versi terbaru.

   ```
   migrate -database "mysql://root:@tcp(127.0.0.1:3306)/eas_pbkk" -path database/migrations up
   ```

   **Down Migration**
   Perintah ini akan membatalkan migrasi terakhir yang diterapkan. Ini berguna ketika Anda ingin mengembalikan skema database ke versi sebelumnya.

   ```
   migrate -database "mysql://root:@tcp(127.0.0.1:3306)/eas_pbkk" -path database/migrations down
   ```

   b. Migration Database With PostgreSQL

   **Up Migration**

   ```
   migrate -database "postgres://postgres:password@127.0.0.1:5432/backendanime?sslmode=disable" -path database/migrations up
   ```

   **Down Migration**

   ```
   migrate -database "postgres://postgres:password@127.0.0.1:5432/backendanime?sslmode=disable" -path database/migrations down
   ```

## Menjalankan Aplikasi

1. Jalankan aplikasi golang:

```sh
go run main.go
```

2. Jalankan services Predict Image Anime

```sh
python services/predict_services.py
```
