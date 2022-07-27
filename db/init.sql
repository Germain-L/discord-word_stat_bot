CREATE DATABASE stats;

CREATE TABLE users(
   id BIGINT,
   name TEXT NOT NULL,
   PRIMARY KEY(id)
);

CREATE TABLE word(
   id SERIAL,
   word_value TEXT,
   PRIMARY KEY(id)
);

CREATE TABLE says(
   user_id BIGINT,
   word_id INT,
   word_count INT,
   FOREIGN KEY(user_id) REFERENCES users(id),
   FOREIGN KEY(word_id) REFERENCES word(id)
);