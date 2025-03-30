CREATE TABLE IF NOT EXISTS `webapp_practice`.`genres` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ジャンルID',
  `name` VARCHAR(64) NOT NULL COMMENT 'ジャンル名',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = 'ジャンル';

CREATE TABLE IF NOT EXISTS `webapp_practice`.`categories` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'カテゴリーID',
  `name` VARCHAR(64) NOT NULL COMMENT 'カテゴリー名',
  `genre_id` INT UNSIGNED NOT NULL COMMENT 'ジャンルID',
  PRIMARY KEY (`id`),
  INDEX `FK_categories_genre_id_idx` (`genre_id` ASC),
  CONSTRAINT `FK_categories_genre_id`
    FOREIGN KEY (`genre_id`)
    REFERENCES `webapp_practice`.`genres` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = 'カテゴリー';

CREATE TABLE IF NOT EXISTS `webapp_practice`.`articles` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '記事ID',
  `title` VARCHAR(128) NOT NULL COMMENT 'タイトル',
  `url` VARCHAR(256) NOT NULL COMMENT '記事URL',
  `published_at` DATETIME NOT NULL COMMENT '記事の公開日時',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = '記事';

CREATE TABLE IF NOT EXISTS `webapp_practice`.`article_categories` (
  `article_id` INT UNSIGNED NOT NULL COMMENT '記事ID',
  `category_id` INT UNSIGNED NOT NULL COMMENT 'カテゴリーID',
  PRIMARY KEY (`article_id`, `category_id`),
  INDEX `FK_article_categories_category_id_idx` (`category_id` ASC),
  CONSTRAINT `FK_article_categories_article_id`
    FOREIGN KEY (`article_id`)
    REFERENCES `webapp_practice`.`articles` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `FK_article_categories_category_id`
    FOREIGN KEY (`category_id`)
    REFERENCES `webapp_practice`.`categories` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = '記事カテゴリーテーブル';

CREATE TABLE IF NOT EXISTS `webapp_practice`.`user_categories` (
  `user_id` INT UNSIGNED NOT NULL COMMENT 'ユーザーID',
  `category_id` INT UNSIGNED NOT NULL COMMENT 'カテゴリーID',
  PRIMARY KEY (`user_id`, `category_id`),
  INDEX `FK_user_categories_category_id_idx` (`category_id` ASC),
  CONSTRAINT `FK_user_categories_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `webapp_practice`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `FK_user_categories_category_id`
    FOREIGN KEY (`category_id`)
    REFERENCES `webapp_practice`.`categories` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = 'ユーザーが関心を持つカテゴリー';

INSERT INTO `webapp_practice`.`genres` (`name`) VALUES
  ('社会'),
  ('経済'),
  ('ニュース'),
  ('スポーツ'),
  ('テクノロジー'),
  ('エンタメ'),
  ('政治'),
  ('科学'),
  ('健康'),
  ('ライフスタイル'),
  ('ビジネス'),
  ('環境'),
  ('教育'),
  ('国際'),
  ('文化');