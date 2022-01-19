CREATE DATABASE IF NOT EXISTS `wager_test`;

USE `wager_test`;
DROP TABLE IF EXISTS `purchases`;
DROP TABLE IF EXISTS `wagers`;
CREATE TABLE IF NOT EXISTS `wagers`
(
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `total_wager_value`      INT(10) DEFAULT 0,
  `odds`      INT(10) DEFAULT 0,
  `selling_percentage`      INT(10) DEFAULT 0,
  `selling_price`      DECIMAL(10, 2) DEFAULT 0,
  `current_selling_price`      DECIMAL(10, 2) DEFAULT 0,
  `percentage_sold`      DECIMAL(10, 2)     DEFAULT NULL,
  `amount_sold`      INT(10)     DEFAULT NULL,
  `placed_at`      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

CREATE TABLE IF NOT EXISTS `purchases`
(
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `wager_id`      INT(10) unsigned NOT NULL,
  `buying_price`      DECIMAL(10, 2),
  `bought_at`      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  CONSTRAINT wager_id FOREIGN KEY (wager_id) REFERENCES wagers (id)
) ENGINE = InnoDB
  DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

INSERT INTO wagers (total_wager_value, odds, selling_percentage, selling_price, current_selling_price, percentage_sold, amount_sold) VALUES(100,1, 10,100.0,99.0, 1.0, 1);
INSERT INTO purchases (wager_id, buying_price) VALUES(1, 1.0);