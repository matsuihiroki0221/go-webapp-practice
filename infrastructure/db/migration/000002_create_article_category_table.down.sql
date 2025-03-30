-- マスターデータ削除
DELETE FROM `webapp_practice`.`genres` WHERE `name` IN 
  ('社会', '経済', 'ニュース', 'スポーツ', 'テクノロジー', 'エンタメ', '政治', 
   '科学', '健康', 'ライフスタイル', 'ビジネス', '環境', '教育', '国際', '文化');

-- 外部キー制約を考慮し、依存関係のあるテーブルを先に削除
DROP TABLE IF EXISTS `webapp_practice`.`user_categories`;
DROP TABLE IF EXISTS `webapp_practice`.`article_categories`;
DROP TABLE IF EXISTS `webapp_practice`.`articles`;
DROP TABLE IF EXISTS `webapp_practice`.`categories`;
DROP TABLE IF EXISTS `webapp_practice`.`genres`;
