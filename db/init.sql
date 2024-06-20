	DROP TABLE IF EXISTS users;
	CREATE TABLE users (
		id SERIAL PRIMARY KEY NOT NULL,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		age INTEGER,
		avatar TEXT,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		password TEXT NOT NULL,
		description TEXT
	);

	DROP TABLE IF EXISTS messages;
	CREATE TABLE messages (
		id SERIAL  PRIMARY KEY NOT NULL,
		content TEXT NOT NULL,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
		status TEXT NOT NULL DEFAULT 'pending',
		fromUserId INTEGER NOT NULL,
		toUserId INTEGER NOT NULL,
    CONSTRAINT fk_from_user_id FOREIGN KEY (fromUserId) REFERENCES users(id),
    CONSTRAINT fk_to_user_id FOREIGN KEY (toUserId) REFERENCES users(id)
	);

INSERT INTO users (name, email, age, avatar, password, description)
VALUES('John Doe', 'jdoe@doemail.com', 23, 'https://cdn.pfps.gg/pfps/9076-amber-rawr.png', '123456', 'A simple todo app built with Go and Tailwind CSS'),
('Bonnie Green', 'bgreen@doemail.com', 23, 'https://i.pinimg.com/736x/5e/19/68/5e196899ff4b98d4352d1ba7337db1ee.jpg', '123456', 'A simple todo app built with Go and Tailwind CSS'),
('John Smith', 'jsmith@doemail.com', 23, 'https://cdn.rudo.video/upload/cl3/adn/640892/5fff0b313d8e7_0.jpg', '123456', 'A simple todo app built with Go and Tailwind CSS');

INSERT INTO messages(content, fromUserId, toUserId, status)
VALUES('Hello', 1, 2, 'delivered'),
('Hi', 2, 1, 'delivered'),
('How are you?', 1, 2, 'delivered'),
('I''m fine', 2, 1, 'delivered'),
('What''s up?', 1, 2, 'delivered'),
('Nothing much', 2, 1, 'delivered'),
('I''m good', 1, 2, 'delivered'),
('Good to see you', 2, 1, 'delivered'),
('Hey', 1, 3, 'pending'),
('How are you?', 3, 1, 'pending'),
('Hey', 3, 2, 'pending');