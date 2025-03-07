CREATE TABLE IF NOT EXISTS `webapp_practice`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT 'ユーザーId',
  `auth0_id` VARCHAR(64) NOT NULL COMMENT 'auth0_id',
  `name` VARCHAR(64) NOT NULL COMMENT 'ユーザー名',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `auth0_id_UNIQUE` (`auth0_id` ASC) VISIBLE
)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = 'ユーザー';

