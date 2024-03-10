CREATE TABLE tasks(
    id int AUTO_INCREMENT PRIMARY KEY,
    user_id int NOT NULL,
    title varchar(255) NOT NULL,
    description text DEFAULT NULL,
    status varchar(50) NOT NULL DEFAULT 'pending',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

