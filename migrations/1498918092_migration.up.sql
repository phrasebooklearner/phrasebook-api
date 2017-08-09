CREATE TABLE users (
  id_user    INT          NOT NULL                 AUTO_INCREMENT PRIMARY KEY,
  name       VARCHAR(255) NOT NULL,
  password   VARCHAR(255) NOT NULL,
  email      VARCHAR(255) NOT NULL,
  created_at TIMESTAMP                             DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY user_email (email(255))
)
  ENGINE = InnoDB