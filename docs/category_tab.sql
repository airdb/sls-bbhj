CREATE TABLE `category_tab` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(1024) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_category_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8

insert into category_tab values(0, null, null, null, "家寻宝贝", "");
insert into category_tab values(0, null, null, null, "宝贝寻家", "");
insert into category_tab values(0, null, null, null, "海外寻亲", "");
insert into category_tab values(0, null, null, null, "其他寻人", "");
insert into category_tab values(0, null, null, null, "流浪乞讨", "");
