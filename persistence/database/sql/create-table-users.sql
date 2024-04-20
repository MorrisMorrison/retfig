CREATE TABLE IF NOT EXISTS user (
    id CHAR(36) NOT NULL, 
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    createdAt DATETIME NOT NULL,
    updatedAt DATETIME NOT NULL,
    isActivated TINYINT(1) DEFAULT 0,
    isEmailConfirmed TINYINT(1) DEFAULT 0,
    PRIMARY KEY (id)
);
