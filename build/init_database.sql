CREATE DATABASE IF NOT EXISTS `wager`;

USE `wager`;
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
  DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC
