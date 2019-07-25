-- Creates a new table for users
CREATE TABLE users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username VARCHAR(64) NOT NULL,
	password VARCHAR(64) NOT NULL,
	UNIQUE(username)
);

-- Creates a new table for messages
CREATE TABLE messages (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	timestamp DATETIME NOT NULL,
	sender INT(11) NOT NULL,
	recipient INT(11) NOT NULL,
	content VARCHAR(255) NOT NULL,
	FOREIGN KEY (sender) REFERENCES users(id),
	FOREIGN KEY (recipient) REFERENCES users(id)
);
