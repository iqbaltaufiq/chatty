-- +goose Up
CREATE TABLE IF NOT EXISTS chats (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    username VARCHAR(32) NOT NULL,
    room_id CHAR(6) NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT UTC_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT UTC_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_chats_users_username
    FOREIGN KEY (username) REFERENCES users (username)
    ON UPDATE cascade
    ON DELETE cascade,
    CONSTRAINT fk_chats_rooms_room_id
    FOREIGN KEY (room_id) REFERENCES rooms (id)
    ON UPDATE cascade
    ON DELETE cascade
) ENGINE = InnoDB;

-- +goose Down
DROP TABLE IF EXISTS chats;
