CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    address VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    born_date TIMESTAMP
    );

-- migrate -database "mysql://root:@tcp(127.0.0.1:3306)/db_gorm" -path database/migrations up
-- migrate -database "mysql://root:@tcp(127.0.0.1:3306)/db_gorm" -path database/migrations down

-- create file migration
-- migrate create -ext sql -dir database/migrations -seq create_users_table