-- +goose Up
CREATE TABLE IF NOT EXISTS occupants (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    room_id CHAR(6) NOT NULL,
    username VARCHAR(32) NOT NULL,
    created_at TIMESTAMP DEFAULT UTC_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_rooms_occupants_room_id
    FOREIGN KEY (room_id) REFERENCES rooms (id)
    ON UPDATE cascade
    ON DELETE cascade,
    CONSTRAINT fk_rooms_occupants_username
    FOREIGN KEY (username) REFERENCES users (username)
    ON UPDATE cascade
    ON DELETE cascade
) ENGINE = InnoDB;

-- +goose Down
DROP TABLE IF EXISTS occupants;
