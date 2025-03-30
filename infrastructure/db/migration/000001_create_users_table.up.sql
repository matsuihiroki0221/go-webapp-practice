CREATE TABLE IF NOT EXISTS `webapp_practice`.`users` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーId',
  `provider_id` VARCHAR(64) NOT NULL COMMENT '外部認証のID（Auth0やGoogleのID）',
  `provider_name` VARCHAR(32) NOT NULL COMMENT '認証プロバイダー名（auth0, google など）',
  `name` VARCHAR(64) NULL COMMENT 'ユーザー名',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_provider` (`provider_id`, `provider_name`) VISIBLE
)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = 'ユーザー';