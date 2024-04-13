CREATE TABLE IF NOT EXISTS event (
    id CHAR(36) NOT NULL, 
    name CHAR(255) NOT NULL,
    creatorEmail CHAR(255) NOT NULL,
    recipient CHAR(255) NOT NULL,
    createdAt DATETIME NOT NULL,
    updatedAt DATETIME NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS event_participant (
    event_id CHAR(36) NOT NULL,
    participant CHAR(255) NOT NULL,
    createdAt DATETIME NOT NULL,
    updatedAt DATETIME NOT NULL,
    PRIMARY KEY (event_id, participant),
    FOREIGN KEY (event_id) REFERENCES event(id)
);
