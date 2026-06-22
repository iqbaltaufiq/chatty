-- +goose Up
CREATE TABLE IF NOT EXISTS relationships (
    id int UNSIGNED NOT NULL AUTO_INCREMENT,
    initiator VARCHAR(32) NOT NULL,
    target VARCHAR(32) NOT NULL,
    created_at TIMESTAMP DEFAULT UTC_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT UTC_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_rels_users_initiator
    FOREIGN KEY (initiator) REFERENCES users (username)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
    CONSTRAINT fk_rels_users_target
    FOREIGN KEY (target) REFERENCES users (username)
    ON UPDATE CASCADE
    ON DELETE CASCADE
) ENGINE = InnoDB;

-- +goose Down
DROP TABLE IF EXISTS relationships;
