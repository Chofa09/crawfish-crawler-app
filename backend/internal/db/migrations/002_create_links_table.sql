DROP TABLE IF EXISTS links;

CREATE TABLE links (
  id INT AUTO_INCREMENT PRIMARY KEY,
  url_id INT,
  href VARCHAR(2048) NOT NULL,
  is_internal BOOLEAN,
  status_code INT,
  is_broken BOOLEAN,
  FOREIGN KEY (url_id) REFERENCES urls(id) ON DELETE CASCADE
);
