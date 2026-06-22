-- +goose Up
-- SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS users (
    uuid BINARY(16) NOT NULL,
    username VARCHAR(32) NOT NULL UNIQUE,
    fullname VARCHAR(128) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    photo VARCHAR(255) NOT NULL DEFAULT "default.png",
    created_at TIMESTAMP DEFAULT UTC_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT UTC_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (uuid)
) ENGINE = InnoDB;

-- +goose Down
-- SELECT 'down SQL query';
DROP TABLE IF EXISTS users;
