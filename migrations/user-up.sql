CREATE DATABASE users;
use users;

CREATE TABLE users(
    id binary(16) not null DEFAULT '',
    first_name longtext,
    last_name longtext,
    user_name varchar(191) NOT NULL,
    email varchar(191) NOT NULL,
    password longblob NOT NULL,
    phone longtext,
    is_admin tinyint(1)
);

INSERT INTO users(id, first_name, last_name, user_name, email, password, phone, is_admin)
VALUES("1", "Rick", "Sanchez", "admin", "rick@c137.com", "$2a$10$Uk2L4qRra2GEluPQdztZqOydkc2pa9DPu.ru2i315jSelVnidrqU2", "5551231234", 1);