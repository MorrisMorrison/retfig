CREATE TABLE events (
    id CHAR(36) NOT NULL, 
    owner CHAR(36) NOT NULL,
    recipient VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (owner) REFERENCES users(id)
);

