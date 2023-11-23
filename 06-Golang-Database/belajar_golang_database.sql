-- Active: 1692450910231@@127.0.0.1@3306@belajar_golang_database
CREATE TABLE customer(
    id VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;
SELECT * FROM customer;
DELETE FROM customer;
ALTER TABLE customer
ADD COLUMN email VARCHAR(100),
    ADD COLUMN balance INT DEFAULT 0,
    ADD COLUMN rating DOUBLE DEFAULT 0.0,
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN birth_date DATE,
    ADD COLUMN married BOOLEAN DEFAULT false;
DESC customer;
INSERT INTO customer(id, name, email, balance, rating, birth_date, married)
VALUES  ("ryu", "Ryu", "ryu@mail.com", 1000000, 90.0, "1999-10-10", true),
        ("jack", "Jack", "jack@mail.com", 500000, 85.5, "1986-02-18", true),
        ("jin", "Jin", "jin@mail.com", 750000, 88.5, "1998-03-18", false);

UPDATE customer
SET email = NULL,
    birth_date = NULL
WHERE id = "jin";

CREATE TABLE user(
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    PRIMARY KEY (username)
)ENGINE = InnoDB;
INSERT INTO user(username, password)
VALUES ('jack', 'jack');
SELECT * FROM user;
CREATE TABLE comments(
    id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(100) NOT NULL,
    comment TEXT,
    PRIMARY KEY (id)
)ENGINE = InnoDB;
DESC comments;
SELECT * FROM comments;
SELECT COUNT(*) FROM comments;