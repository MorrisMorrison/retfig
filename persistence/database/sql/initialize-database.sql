CREATE TABLE IF NOT EXISTS event (
    id CHAR(36) NOT NULL, 
    name CHAR(255) NOT NULL,
    recipient CHAR(255) NOT NULL,
    createdBy CHAR(255) NOT NULL,
    updatedBy CHAR(255) NOT NULL,
    createdAt DATETIME NOT NULL,
    updatedAt DATETIME NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS event_participant (
    eventId CHAR(36) NOT NULL,
    name CHAR(255) NOT NULL,
    createdBy CHAR(255) NOT NULL,
    updatedBy CHAR(255) NOT NULL,
    createdAt DATETIME NOT NULL,
    updatedAt DATETIME NOT NULL,
    PRIMARY KEY (eventId, name),
    FOREIGN KEY (eventId) REFERENCES event(id)
);

CREATE TABLE IF NOT EXISTS present ( 
    id CHAR(36) NOT NULL, 
    eventId CHAR(36) NOT NULL,
    name CHAR(255) NOT NULL,
    link CHAR(255) NOT NULL,
    createdBy CHAR(255) NOT NULL,
    updatedBy CHAR(255) NOT NULL,
    createdAt DATETIME NOT NULL,
    updatedAt DATETIME NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (eventId) REFERENCES event(id)
);

CREATE TABLE IF NOT EXISTS present_vote (
    presentId CHAR(36) NOT NULL,
    type CHAR(36) NOT NULL, 
    createdBy CHAR(255) NOT NULL,
    updatedBy CHAR(255) NOT NULL,
    createdAt DATETIME NOT NULL,
    updatedAt DATETIME NOT NULL,
    PRIMARY KEY (presentId, createdBy),
    FOREIGN KEY (presentId) REFERENCES present(id)
);

CREATE TABLE IF NOT EXISTS present_comment (
    presentId CHAR(36) NOT NULL,
    content CHAR(255) NOT NULL, 
    createdBy CHAR(255) NOT NULL,
    updatedBy CHAR(255) NOT NULL,
    createdAt DATETIME NOT NULL,
    updatedAt DATETIME NOT NULL,
    FOREIGN KEY (presentId) REFERENCES present(id)
);