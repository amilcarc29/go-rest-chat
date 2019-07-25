-- Creates a new table for users
CREATE TABLE users (
	id INT(11) NOT NULL AUTO_INCREMENT,
	username VARCHAR(64) NOT NULL,
	password VARCHAR(64) NOT NULL,
	created_at DATETIME NULL,
	updated_at DATETIME NULL,
	deleted_at DATETIME NULL,
	PRIMARY KEY(id),
	UNIQUE KEY (username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_unicode_ci;

-- Creates a new table for the content of a message
CREATE TABLE contents (
	id INT(11) NOT NULL AUTO_INCREMENT,
	type ENUM('text', 'image', 'video') NOT NULL,
	text VARCHAR(255) NOT NULL,
	created_at DATETIME NULL,
	updated_at DATETIME NULL,
	deleted_at DATETIME NULL,
	PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_unicode_ci;

-- Creates a new table for messages
CREATE TABLE messages (
	id INT(11) NOT NULL AUTO_INCREMENT,
	sender_id INT(11) NOT NULL,
	recipient_id INT(11) NOT NULL,
	content_id INT(11) NOT NULL,
	created_at DATETIME NULL,
	updated_at DATETIME NULL,
	deleted_at DATETIME NULL,
	PRIMARY KEY(id),
	FOREIGN KEY (sender_id) REFERENCES users(id),
	FOREIGN KEY (recipient_id) REFERENCES users(id),
	FOREIGN KEY (content_id) REFERENCES contents(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_unicode_ci;
