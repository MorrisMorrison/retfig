CREATE TABLE users (
    id CHAR(36) NOT NULL, 
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    createdAt DATETIME NOT NULL,
    createdBy CHAR(36) NOT NULL,
    updatedAt DATETIME NOT NULL,
    updatedBy CHAR(36) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (createdBy) REFERENCES users(id),
    FOREIGN KEY (updatedBy) REFERENCES users(id)
);

