DROP TABLE IF EXISTS urls;

CREATE TABLE urls (
  id INT AUTO_INCREMENT PRIMARY KEY,
  address VARCHAR(2048) NOT NULL,
  page_title VARCHAR(500),
  html_version VARCHAR(20),
  headings JSON,
  has_login_form BOOLEAN,
  crawl_status ENUM('queued', 'running', 'done', 'error'),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
