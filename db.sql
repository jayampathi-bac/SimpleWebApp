
drop database simplecruddb;
create database if not exists simplecruddb;

use simplecruddb;

CREATE TABLE IF NOT EXISTS user (
  id VARCHAR(20) NOT NULL,
  username VARCHAR(100) NULL,
  email VARCHAR(450) NULL,
  address VARCHAR(450) NULL,
  CONSTRAINT PRIMARY KEY (id)
);