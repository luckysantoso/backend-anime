CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    address VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    born_date TIMESTAMP
    );

-- CREATE TABLE users (
--     id AUTO_INCREMENT PRIMARY KEY,
--     name VARCHAR(255),
--     address VARCHAR(255),
--     email VARCHAR(255) NOT NULL,
--     born_date TIMESTAMP
--     );

-- MySQL
-- migrate -database "mysql://root:@tcp(127.0.0.1:3306)/db_gorm" -path database/migrations up
-- migrate -database "mysql://root:@tcp(127.0.0.1:3306)/db_gorm" -path database/migrations down

-- Postgres
-- migrate -database "postgres://postgres:Subhanallah25@127.0.0.1:5432/backendanime?sslmode=disable" -path database/migrations up
-- migrate -database "postgres://postgres:Subhanallah25@127.0.0.1:5432/backendanime?sslmode=disable" -path database/migrations down
-- migrate -database "postgres://postgres:Subhanallah25@127.0.0.1:5432/backendanime?sslmode=disable" -path database/migrations force 1


-- create file migration
-- migrate create -ext sql -dir database/migrations -seq create_users_table