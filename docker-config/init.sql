CREATE DATABASE IF NOT EXISTS company;
USE company;
CREATE TABLE employee (
    id INT AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    PRIMARY KEY(id)
) ENGINE=InnoDB;
INSERT INTO employee (name) VALUES ("Eduardo Test 1");