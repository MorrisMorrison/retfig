CREATE TABLE event_participants (
    event_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    PRIMARY KEY (event_id, user_id),
    FOREIGN KEY (event_id) REFERENCES events(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

