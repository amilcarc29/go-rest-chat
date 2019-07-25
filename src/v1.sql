-- Creates a new table for users
CREATE TABLE users (
	id INT(11) NOT NULL AUTO_INCREMENT,
	username VARCHAR(64) NOT NULL,
	password VARCHAR(64) NOT NULL,
	PRIMARY KEY(id),
	UNIQUE KEY (username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_unicode_ci;

-- Creates a new table for messages
CREATE TABLE messages (
	id INT(11) NOT NULL AUTO_INCREMENT,
	timestamp DATETIME NOT NULL,
	sender INT(11) NOT NULL,
	recipient INT(11) NOT NULL,
	content VARCHAR(255) NOT NULL,
	PRIMARY KEY(id),
	FOREIGN KEY (sender) REFERENCES users(id),
	FOREIGN KEY (recipient) REFERENCES users(id),
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_unicode_ci;
