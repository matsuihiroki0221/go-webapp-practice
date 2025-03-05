
CREATE TABLE IF NOT EXISTS `webapp_practice`.`categories` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` INT NOT NULL COMMENT 'ユーザーID',
  `name` VARCHAR(64) NOT NULL COMMENT 'カテゴリー名',
  PRIMARY KEY (`id`),
  INDEX `FK_categories_user_id_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `FK_categories_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `webapp_practice`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
COMMENT = 'カテゴリー';

CREATE TABLE IF NOT EXISTS `webapp_practice`.`articles` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '記事ID',
  `user_id` INT NOT NULL COMMENT 'ユーザID',
  `title` VARCHAR(128) NOT NULL COMMENT 'タイトル',
  `url` VARCHAR(256) NOT NULL COMMENT '記事URL',
  PRIMARY KEY (`id`),
  INDEX `FK_articles_user_id_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `FK_articles_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `webapp_practice`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = '記事';

CREATE TABLE IF NOT EXISTS `webapp_practice`.`article_categories` (
  `article_id` INT NOT NULL COMMENT '記事ID',
  `category_id` INT NOT NULL COMMENT 'カテゴリーID',
  PRIMARY KEY (`article_id`, `category_id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = '記事カテゴリーテーブル';
