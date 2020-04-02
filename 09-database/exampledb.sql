CREATE DATABASE gotest;

CREATE TABLE IF NOT EXISTS tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    start_date DATE,
    due_date DATE,
    status TINYINT NOT NULL,
    priority TINYINT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)  ENGINE=INNODB;

INSERT INTO tasks (title,start_date,due_date,status,priority,description,created_at) VALUES ('Go ile Test Yazmayı Öğren',now(),now() + interval 1 month,1,1,"Test yazmak gerçekten çok önemlidir. Kodu kalitesini artırır. Hataları giderir.",now());