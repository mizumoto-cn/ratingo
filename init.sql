CREATE TABLE Users (
  id INT AUTO_INCREMENT,
  name VARCHAR(255),
  email VARCHAR(255) UNIQUE,
  PRIMARY KEY (id)
);

CREATE TABLE Topics (
  id INT AUTO_INCREMENT,
  topic_name VARCHAR(255),
  PRIMARY KEY (id)
);

CREATE TABLE Ratings (
  id INT AUTO_INCREMENT,
  user_id INT,
  topic_id INT,
  rating FLOAT,
  comment TEXT,
  FOREIGN KEY (user_id) REFERENCES Users(id),
  FOREIGN KEY (topic_id) REFERENCES Topics(id),
  PRIMARY KEY (id)
);

CREATE TABLE Analysis (
  id INT AUTO_INCREMENT,
  topic_id INT,
  avg_rating FLOAT,
  FOREIGN KEY (topic_id) REFERENCES Topics(id),
  PRIMARY KEY (id)
);
