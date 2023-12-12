CREATE TABLE users (
  id INT AUTO_INCREMENT,
  name VARCHAR(255),
  email VARCHAR(255) UNIQUE,
  PRIMARY KEY (id)
);

CREATE TABLE topics (
  id INT AUTO_INCREMENT,
  topic_name VARCHAR(255),
  PRIMARY KEY (id)
);

CREATE TABLE ratings (
  id INT AUTO_INCREMENT,
  user_id INT,
  topic_id INT,
  rating FLOAT,
  comment TEXT,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (topic_id) REFERENCES topics(id),
  PRIMARY KEY (id)
);

CREATE TABLE analyses (
  id INT AUTO_INCREMENT,
  topic_id INT,
  avg_rating FLOAT,
  FOREIGN KEY (topic_id) REFERENCES topics(id),
  PRIMARY KEY (id)
);
