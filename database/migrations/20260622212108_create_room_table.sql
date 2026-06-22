-- +goose Up
CREATE TABLE IF NOT EXISTS rooms (
    id CHAR(6),
    name VARCHAR(64) NOT NULL,
    is_private BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT UTC_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE = InnoDB;

-- +goose Down
DROP TABLE IF EXISTS rooms;
