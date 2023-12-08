CREATE DATABASE IF NOT EXISTS backend;
USE backend;

CREATE TABLE IF NOT EXISTS user (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255),
    password VARCHAR(255),
    create_time DATETIME,
    update_time DATETIME,
    avatar TEXT
);