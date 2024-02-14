CREATE TABLE IF NOT EXISTS stored_user (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255),
    password VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS security (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS stored_order (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    price INT,
    quantity INT,
    type INT,
    fulfilled BOOLEAN DEFAULT FALSE NOT NULL,
    security_id BIGINT,
    user_id BIGINT
);

INSERT INTO stored_user (id, username, password) VALUES
(1, 'user1', 'password1'),
(2, 'user2', 'password2'),
(3, 'user3', 'password3'),
(4, 'user4', 'password4');

INSERT INTO security (id, name) VALUES
(1, 'Apple'),
(2, 'Google');

INSERT INTO stored_order (price, quantity, type, security_id, user_id) VALUES
(100, 50, 0, 1, 1),
(200, 100, 1, 2, 2);
